package main

import (
	"fmt"
	"log"
	"net/http"
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

func main() {
	router := NewRouter()

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	fmt.Printf("Server running on port 4000 \n")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
