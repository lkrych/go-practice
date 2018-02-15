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

func handleAnswerCreate(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	//declaring an anonymous struct that embeds answer, when we decode this
	//it will capture Answer fields as well as QuestionID
	var newAnswer struct {
		Answer
		QuestionID string `json:"question_id"`
	}

	err := decode(r, &newAnswer)
	if err != nil {
		respondErr(ctx, w, r, err, http.StatusBadRequest)
		return
	}

	questionKey, err := datastore.DecodeKey(newAnswer.QuestionID)
	if err != nil {
		respondErr(ctx, w, r, err, http.StatusBadRequest)
		return
	}

	err = newAnswer.OK()
	if err != nil {
		respondErr(ctx, w, r, err, http.StatusBadRequest)
		return
	}

	answer := newAnswer.Answer
	user, err := UserFromAEUser(ctx)
	if err != nil {
		respondErr(ctx, w, r, err, http.StatusBadRequest)
		return
	}

	answer.User = user.Card()
	err = answer.Create(ctx, questionKey)
	if err != nil {
		respondErr(ctx, w, r, err, http.StatusInternalServerError)
		return
	}

	respond(ctx, w, r, answer, http.StatusCreated)
	}
}
