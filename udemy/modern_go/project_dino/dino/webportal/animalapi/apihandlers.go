package animalapi

import (
	"dino/databaselayer"
)

type AnimalRESTAPIHandler struct {
	dbhandler databaselayer.AnimalDBHandler
}

func newAnimalRESTAPIHandler(db databaselayer.AnimalDBHandler) *AnimalRESTAPIHandler {
	return &AnimalRESTAPIHandler{
		dbhandler: db,
	}
}

func (handler *AnimalRESTAPIHandler) searchHandler
