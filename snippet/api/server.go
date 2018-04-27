package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Handle API routes
	api := r.PathPrefix("/api/").Subrouter()

	api.Handle("/search/{keyword}", searchTermHandler).Methods("GET")

	// Serve static assets directly.
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	fmt.Printf("Serving api on port 8080 \n")
	http.ListenAndServe(":8080", r)
}