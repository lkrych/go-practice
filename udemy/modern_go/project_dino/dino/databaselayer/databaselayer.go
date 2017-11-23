package databaselayer

import (
	"errors"
)

const ( //this is how you create an enum
	MYSQL uint8 = iota
	SQLITE
	POSTGRESSQL
	MONGODB
)

type AnimalDBHandler interface {
	GetAvailableAnimals() ([]Animal, error)
	GetAnimalByNickname(string) (Animal, error)
	GetAnimalsByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

var DBTypeNotSupported = errors.New("The Database type provided is not supported...")

//factory function
func GetDatabaseHandler(dbtype uint8, connection string) (AnimalDBHandler, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongdbHandler(connection)
	case POSTGRESSQL:
		return NewPGHandler(connection)
	case SQLITE:
		return NewSQLiteHandler(connection)
	}
	return nil, DBTypeNotSupported
}
