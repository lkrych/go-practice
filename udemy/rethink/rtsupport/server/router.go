package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // overwrites default behavior of library, allows connections from any origin
}

type Router struct {
	rules map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerErrror)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(socket)
	go client.Write()
	client.Read()
}
