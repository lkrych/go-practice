package main

import (
	"log"
	"net/http"
	"trace"

	"github.com/gorilla/websocket"
)

//note: channels are an in-memory thread-safe message queue where
// senders pass data and receivers read data in a non-blocking thread-safe way.

//the practical significance of this is that if we tried to directly access the clients map
// it is possible that two goroutines running concurently might try to modify the map at the same time,
//resulting in corrupt memory or unpredictable state

type room struct {
	//channel that holds incoming messages that should be forwarded to other clients
	forward chan []byte
	//join is a channel for clients wishing to join the room
	join chan *client
	//leave is a channel for clients wishing to leave the room
	leave chan *client
	//client holds all current clients in this room
	clients map[*client]bool
	// tracer will receive trace information activity
	tracer trace.Tracer
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

// the select statement can be used whenever we need
// to synchronize or modify shared memory

// if this function is run in a goroutine, it will run in the background
// which won't block the rest of our application. This code will watch the three channels inside
// the room and will run the case statement if a msg is received

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//joining
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			//leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", string(msg))
			//forward message to all clients
			for client := range r.clients {
				//msg will be sent into client's send channel
				//client's write method will pick it up
				client.send <- msg
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 1024
)

//make room type into an http.Handler type
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read() // will block main thread until it's time to close it
}
