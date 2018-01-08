package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	r "gopkg.in/gorethink/gorethink.v4"
)

type Channel struct {
	ID   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"data" gorethink:"name"`
}

type User struct {
	ID   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

type ChannelMessage struct {
	ID        string    `gorethink:"id,omitempty"`
	ChannelID string    `gorethink:"channelID"`
	Body      string    `gorethink:"body"`
	Author    string    `gorethink:"author"`
	CreatedAt time.Time `gorethink:"createdAt"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})

	if err != nil {
		log.Panic(err.Error())
	}
	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)

	router.Handle("user edit", editUser)
	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)

	router.Handle("message add", addChannelMesage)
	router.Handle("message subscribe", subscribeChannelMessage)
	router.Handle("message unsubscribe", unsubscribeChannelMessage)

	http.Handle("/", router)
	fmt.Printf("Server running on port 4000 \n")
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
