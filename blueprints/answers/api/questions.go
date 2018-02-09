package api

import (
	"context"
	"errors"
	"time"

	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

type Question struct {
	Key          *datastore.Key `json: "id" datastore: "-"`
	CTime        time.Time      `json: "created"`
	Question     string         `json: "question"`
	User         UserCard       `json: "user"`
	AnswersCount int            `json: "answers_count"`
}

//OK checks to make sure that question length is valid
func (q Question) OK() error {
	if len(q.Question) < 10 {
		return errors.New("question is too short")
	}
	return nil
}

//Create builds a Question in the Datastore (a NoSQL db)
func (q *Question) Create(ctx context.Context) error {
	log.Debugf(ctx, "Saving question: %s", q.Question)
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}
	user, err := UserFromAEUser(ctx)
	if err != nil {
		return err
	}
	q.User = user.Card()
	q.CTime = time.Now()
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}
	return nil
}

//Update updates a question in the Datastore
func (q *Question) Update(ctx context.Context) error {
	if q.Key == nil {
		q.Key = datastore.NewIncompleteKey(ctx, "Question", nil)
	}
	var err error
	q.Key, err = datastore.Put(ctx, q.Key, q)
	if err != nil {
		return err
	}
	return nil
}

//GetQuestion returns a Question struct with the key specified in the argument
func GetQuestion(ctx context.Context, key *datastore.Key) (*Question, error) {
	var q Question
	err := datastore.Get(ctx, key, &q)
	if err != nil {
		return nil, err
	}
	q.Key = key
	return &q, nil
}
