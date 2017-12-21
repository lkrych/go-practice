package main

//note: channels are an in-memory thread-safe message queue where
// senders pass data and receivers read data in a non-blocking thread-safe way.

type room struct {
	//channel that holds incoming messages that should be forwarded to other clients
	forward chan []byte
	//join is a channel for clients wishing to join the room
	join chan *client
	//leave is a channel for clients wishing to leave the room
	leave chan *client
	//client holds all current clients in this room
	clients map[*client]bool
}
