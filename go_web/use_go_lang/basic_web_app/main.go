package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to this awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+"to <a href=\"mailto:support@lenslocked.com\">"+"support@lenslocked.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<ul><li>Q: How do I get home?</li> <li>A: Search for the URL http://localhost:3000/ or click <a href=\"http://localhost:3000/\">here</a></li></ul>")
}

func fourOhfour(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>🌼 Whoopsy Daisy 🌼</h1><p>This isn't what you were looking for! Get on back <a href=\"http://localhost:3000/\">home</a>!</p>")
}

func main() {
	var h http.Handler = http.HandlerFunc(fourOhfour)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}
