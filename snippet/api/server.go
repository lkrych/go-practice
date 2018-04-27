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

	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(IndexHandler("../frontend/index.html"))
	fmt.Printf("Serving api on port 3000 \n")
	http.ListenAndServe(":3000", r)
}
