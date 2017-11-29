package main

import (
	"dino/databaselayer"
	"dino/webportal/animalapi"
	"log"
)

func main() {
	db, err := databaselayer.GetDatabaseHandler(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	animalapi.RunAPI("localhost:8080", db)
}
