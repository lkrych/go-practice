package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

//Message is how json will be received in the server, we use an empty interface to accept any type of data
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"data"`
}

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
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)
		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage = Message{"error", err}
				if err := socket.WriteJSON(outMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			//make sure subscribe doesn't block
			go subscribeChannel(socket)
		}
		// fmt.Println(string(msg))
		// if err = socket.WriteMessage(msgType, msg); err != nil { //echo back to client
		// 	log.Fatal(err)
		// }
	}
}

func addChannel(data interface{}) error {
	var channel Channel

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.ID = "1"
	fmt.Printf("%#v\n", channel)
	return nil
}

func subscribeChannel(socket *websocket.Conn) {
	//TODO: rethinkDB query / changefeed
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add", Channel{
			"1", "Software Support"}}
		socket.WriteJSON(message)
	}
}
