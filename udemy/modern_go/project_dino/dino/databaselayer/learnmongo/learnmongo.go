package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	// animals := []interface{}{animal{
	// 	AnimalType: "Beluga Whale",
	// 	Nickname:   "Whitey",
	// 	Zone:       1,
	// 	Age:        5,
	// }, animal{
	// 	AnimalType: "Orangutan",
	// 	Nickname:   "Orange Crush",
	// 	Zone:       2,
	// 	Age:        17,
	// }, animal{
	// 	AnimalType: "CuttleFish",
	// 	Nickname:   "Cuddly",
	// 	Zone:       3,
	// 	Age:        2,
	// },
	// }

	//insert

	// err = animalcollection.Insert(animals...)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//update

	// err = animalcollection.Update(bson.M{"nickname": "Cuddly"}, bson.M{"$set": bson.M{"age": 3}})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//remove

	// err = animalcollection.Remove(bson.M{"nickname": "Cuddly"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//query

	query := bson.M{
		"age": bson.M{
			"$gt": 3,
		},
		"zone": bson.M{
			"$in": []int{1, 2},
		},
	}
	results := []animal{}
	animalcollection.Find(query).All(&results)
	fmt.Println(results)

	//one result
	result := animal{}
	animalcollection.Find(query).One(&result)
	fmt.Println(result)
}
