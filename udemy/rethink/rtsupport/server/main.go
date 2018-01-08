package main

import (
	"fmt"
	"log"
	"net/http"

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

	http.Handle("/", router)
	fmt.Printf("Server running on port 4000 \n")
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
