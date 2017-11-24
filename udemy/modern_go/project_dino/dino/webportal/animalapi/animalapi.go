package animalapi

import (
	"dino/databaselayer"
	"net/http"

	"github.com/gorilla/mux"
)

func RunAPI(endpoint string, db databaselayer.AnimalDBHandler) error {
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router, db databaselayer.AnimalDBHandler) {
	handler := newAnimalRESTAPIHandler(db)

	apirouter := r.PathPrefix("/api/dinos").Subrouter()
	apirouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.searchHandler)
	apirouter.Methods("POST").PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)

}
