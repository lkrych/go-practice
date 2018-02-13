package api

import (
	"context"
	"time"

	"google.golang.org/appengine/datastore"
)

type Vote struct {
	Key      *datastore.Key `json:"id" datastore:"-"`
	MTime    time.Time      `json:"last_modified" datastore:",noindex"`
	Question QuestionCard   `json:"question" datastore:",noindex"`
	Answer   AnswerCard     `json:"answer" datastore:",noindex"`
	User     UserCard       `json:"user" datastore:",noindex"`
	Score    int            `json:"score" datastore:",noindex"`
}

func CastVote(ctx context.Context, answerKey *datastore.Key, score int) (*Vote, error) {
	question, err := GetQuestion(ctx, answerKey.Parent())
	if err != nil {
		return nil, err
	}
	user, err := UserFromAEUser(ctx)
	if err != nil {
		return nil, err
	}
	var vote Vote
	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		var err error
		vote, err = castVoteInTransaction(ctx, answerKey, question, user, score)
		if err != nil {
			return err
		}
		return nil
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		return nil, err
	}
	return &vote, nil
}
