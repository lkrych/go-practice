package api

import datastore "google.golang.org/genproto/googleapis/datastore/v1"

type User struct {
	Key         *datastore.Key `json:"id" datastore:"-"`
	UserID      string         `json:"-"`
	DisplayName string         `json:"display_name"`
	AvatarURL   string         `json:"avatar_url"`
	Score       int            `json:"score"`
}
