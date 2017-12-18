package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // overwrites default behavior of library, allows connections from any origin
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Server running on port 4000 \n")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil { //echo back to client
			log.Fatal(err)
		}
	}
}
