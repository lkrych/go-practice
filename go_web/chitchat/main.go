package main

import (
	"net/http"
)

func main() {

	//init mux and handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	//all route patterns matched here
	//defined in route files
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
