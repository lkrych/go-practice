package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	InitKeys()
	r := mux.NewRouter()

	//Public endpoint
	r.HandleFunc("/login", LoginHandler)

	// Handle API routes
	api := r.PathPrefix("/api/").Subrouter()

	api.Handle("/search/{keyword}", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(searchTermHandler)),
	))

	// Serve static assets directly.
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	fmt.Printf("Serving api on port 8080 \n")
	http.ListenAndServe(":8080", r)
}
