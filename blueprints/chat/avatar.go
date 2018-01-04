package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
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
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return fmt.Sprintf("//www.gravatar.com/avatar/%s", useridStr), nil
		}
	}
	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatarURL
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := path.Match(useridStr+"*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
