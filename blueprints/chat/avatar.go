package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

//ErrNoAvatar is the error that is returned when the Avatar instance is unable to provide an avatar URL
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

//Avatar represents types capable of representing user profile pictures
type Avatar interface {
	//GetAvatarURL gets the avatar URL for the specified client, or returns an error if something goes wrong.
	//ErrNoAvatarURL is returned if the object is unable to get a URL for the specified client.
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar // variable with AuthAvatar type but nil value, can later assign this variable to any field looking for an Avatar interface type

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	urlStr, ok := c.userData["avatar_url"].(string)
	if !ok {
		return "", ErrNoAvatarURL
	}
	return urlStr, nil
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatarURL
}
