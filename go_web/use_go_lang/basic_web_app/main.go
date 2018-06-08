package main

import (
	"net/http"

	"flatphoto.com/controllers"
	"flatphoto.com/views"
	"github.com/gorilla/mux"
)

//global variables are usually frowned upon b/c they make code harder to test and can have side effects
var (
	fourOhfourView *views.View
)

func fourOhfour(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	must(fourOhfourView.Render(w, nil))
}

func main() {
	fourOhfourView = views.NewView("bootstrap", "views/404.gohtml")

	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()

	var h http.Handler = http.HandlerFunc(fourOhfour)
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}

// A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}
