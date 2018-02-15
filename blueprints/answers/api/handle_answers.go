package api

import "net/http"

func handleAnswers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleAnswersGet(w, r)
	case "POST":
		handleAnswersCreate(w, r)
	default:
		http.NotFound(w, r)
	}
}
