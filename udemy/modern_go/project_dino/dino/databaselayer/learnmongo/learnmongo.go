package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type animal struct {
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//get a collection
	animalcollection := session.DB("Dino").C("animals")

	animals := []interface{}{animal{
		AnimalType: "Beluga Whale",
		Nickname:   "Whitey",
		Zone:       1,
		Age:        5,
	}, animal{
		AnimalType: "Orangutan",
		Nickname:   "Orange Crush",
		Zone:       2,
		Age:        17,
	}, animal{
		AnimalType: "CuttleFish",
		Nickname:   "Cuddly",
		Zone:       3,
		Age:        2,
	},
	}
	err = animalcollection.Insert(animals...)
	if err != nil {
		log.Fatal(err)
	}
}
