package animalapi

import (
	"dino/databaselayer"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type AnimalRESTAPIHandler struct {
	dbhandler databaselayer.AnimalDBHandler
}

func newAnimalRESTAPIHandler(db databaselayer.AnimalDBHandler) *AnimalRESTAPIHandler {
	return &AnimalRESTAPIHandler{
		dbhandler: db,
	}
}

func (handler *AnimalRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No search criteria found. You can search by nickname or by type.")
		return
	}
	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No search criteria found. You can search by nickname or by type.")
		return
	}
	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error
	switch strings.ToLower(criteria) {
	case "nickname":
		animal, err = handler.dbhandler.GetAnimalByNickname(searchkey)
	case "type":
		animals, err = handler.dbhandler.GetAnimalsByType(searchkey)
		if len(animals) > 0 {
			json.NewEncoder(w).Encode(animals)
			return
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error occurred while querying animals %v", err)
	}
	json.NewEncoder(w).Encode(animal)
}

func (handler *AnimalRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operation, ok := vars["Operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No Operation found. You can use /api/animals/add to add a new animal, or /api/dinos/edit/:nickname to edit an existing animal.")
		return
	}
	var animal databaselayer.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not decode the request body to json %v", err)
		return
	}
	switch strings.ToLower(operation) {
	case "add":
		err = handler.dbhandler.AddAnimal(animal)
	case "edit":
		nickname := r.RequestURI[len("api/animals/edit/"):]
		log.Println("edit requested for nickname", nickname)
		err = handler.dbhandler.UpdateAnimal(animal, nickname)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error occurred while processing request %v", err)
	}
}
