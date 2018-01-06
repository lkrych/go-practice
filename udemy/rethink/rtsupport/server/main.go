package main

import (
	"fmt"
	"log"
	"net/http"
)

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
