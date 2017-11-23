package databaselayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongdbHandler(connection string) (*MongodbHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: s,
	}, err
}

func (handler *MongodbHandler) GetAvailableAnimals() ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(nil).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) GetAnimalByNickname(nickname string) (Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	a := Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"nickname": nickname}).One(&a)
	return a, err
}

func (handler *MongodbHandler) GetAnimalsByType(animalType string) ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"animal_type": animalType}).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) AddAnimal(a Animal) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Insert(a)
}

func (handler *MongodbHandler) UpdateAnimal(a Animal, nname string) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Update(bson.M{"nickname": nname}, a)
}

//obtains another session to the same database from a connection pool
func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
