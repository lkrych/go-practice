package main

import (
	"fmt"
	"net/http"

	"flatphoto.com/controllers"
	"flatphoto.com/models"
	"flatphoto.com/views"
	"github.com/gorilla/mux"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "flatphoto_dev"
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
	//Create a DB connection string and then use it to create our model services
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

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
