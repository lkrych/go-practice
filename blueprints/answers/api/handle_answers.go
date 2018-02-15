package api

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

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

//handleAnswersGet uses url query parameters, not path parameters, this is
//to demonstrate different styles of api design, ordinarily you should be
//consistent
func handleAnswersGet(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	q := r.URL.Query()
	questionIDStr := q.Get("question_id")
	questionKey, err := datastore.DecodeKey(questionIDStr)
	if err != nil {
		respondErr(ctx, w, r, err, http.StatusBadRequest)
		return
	}
	answers, err := GetAnswers(ctx, questionKey)
	if err != nil {
		respondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}
	respond(ctx, w, r, answers, http.StatusOK)
}
