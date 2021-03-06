package api

import (
	"errors"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

type Answer struct {
	Key    *datastore.Key `json:"id" datastore:"-"`
	Answer string         `json:"answer" datastore:",noindex"`
	CTime  time.Time      `json:"created"`
	User   UserCard       `json:"user" datastore:",noindex"`
	Score  int            `json:"score"`
}

func (a Answer) OK() error {
	if len(a.Answer) < 10 {
		return errors.New("answer is too short")
	}
	return nil
}

//Create updates answer and question using datastore Transaction
func (a *Answer) Create(ctx context.Context, questionKey *datastore.Key) error {
	a.Key = datastore.NewIncompleteKey(ctx, "Answer", questionKey)
	user, err := UserFromAEUser(ctx)
	if err != nil {
		return err
	}
	a.User = user.Card()
	a.CTime = time.Now()
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		q, err := GetQuestion(ctx, questionKey)
		if err != nil {
			return err
		}
		err = a.Put(ctx)
		if err != nil {
			return err
		}
		q.AnswersCount++
		err = q.Update(ctx)
		if err != nil {
			return err
		}
		return nil
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		return err
	}
	return nil
}

//GetAnswer returns an Answer struct
func GetAnswer(ctx context.Context, answerKey *datastore.Key) (*Answer, error) {
	var answer Answer
	err := datastore.Get(ctx, answerKey, &answer)
	if err != nil {
		return nil, err
	}
	answer.Key = answerKey
	return &answer, nil
}

//Put is a method for abstracting updating an Answer
func (a *Answer) Put(ctx context.Context) error {
	var err error
	a.Key, err = datastore.Put(ctx, a.Key, a)
	if err != nil {
		return err
	}
	return nil
}

//GetAnswers fetches answers from the datastore ordered by score and date
func GetAnswers(ctx context.Context, questionKey *datastore.Key) ([]*Answer, error) {
	var answers []*Answer
	answerKeys, err := datastore.NewQuery("Answer").
		Ancestor(questionKey).
		Order("-Score").
		Order("-CTime").
		GetAll(ctx, &answers)
	for i, answer := range answers {
		answer.Key = answerKeys[i]
	}
	if err != nil {
		return nil, err
	}
	return answers, nil
}

//AnswerCard is the friendly version of the Answer struct
type AnswerCard struct {
	Key    *datastore.Key `json:"id" datastore:",noindex"`
	Answer string         `json:"answer" datastore:",noindex"`
	User   UserCard       `json:"user" datastore:",noindex"`
}

//Card is the AnswerCard constructor
func (a Answer) Card() AnswerCard {
	return AnswerCard{
		Key:    a.Key,
		Answer: a.Answer,
		User:   a.User,
	}
}
