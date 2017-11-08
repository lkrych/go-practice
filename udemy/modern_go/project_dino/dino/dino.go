package main

import (
	"dino/server"
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	ServerAddress string `json:"webserver"` //whenever you create structs that you parse from JSON strings, the fields from the structs need to be exportable
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)         //creates a pointer of configuration struct
	json.NewDecoder(file).Decode(config) //read json file and write it to config structure
	log.Println("Starting web server on address ", config.ServerAddress)
	server.RunServer(config.ServerAddress)
}
