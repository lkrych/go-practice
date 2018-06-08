package main

import (
	"net/http"

	"flatphoto.com/controllers"
	"flatphoto.com/views"
	"github.com/gorilla/mux"
)

//global variables are usually frowned upon b/c they make code harder to test and can have side effects
var (
	homeView       *views.View
	contactView    *views.View
	faqView        *views.View
	fourOhfourView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func fourOhfour(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	must(fourOhfourView.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	fourOhfourView = views.NewView("bootstrap", "views/404.gohtml")

	usersC := controllers.NewUsers()

	var h http.Handler = http.HandlerFunc(fourOhfour)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/signup", usersC.New)
	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}

// A helper function that panics on any error
func must(err error) {
	if err != nil {
		panic(err)
	}
}
