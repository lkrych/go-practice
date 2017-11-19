package databaselayer

import (
	"errors"
)

const { //this is how you create an enum
	MYSQL uint8 - iota
	SQLITE
	POSTGRESSQL
	MONGODB
}

type DinoDBHandler interface {
	GetAvailableAnimals() ([]Animal, error)
	GetAnimalsByNickname(string) (Animal, error)
	GetAnimalByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {	
	ID int `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname string `bson:"nickname"`
	Zone int `bson:"zone"`
	Age int `bson:"age"`
}

var DBTypeNotSupported = errors.New("The Database type provided is not supported...")

//factory function
func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLHANDLER(connection)
	case MONGODB:
		return NewMongodbHandler(connection)
	case POSTGRESQL:
		return NewPGHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	}
	return nil, DBTypeNotSupported
}