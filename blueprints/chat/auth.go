package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")

	if err == http.ErrNoCookie {
		//not authenticated
		w.Header().Set("Location", "/login")

		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//if you don't run into any errors, call the next handler

	h.next.ServeHTTP(w, r)
}

// MustAuth creates an auth handler struct that wraps any other http handler
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

//loginHandler handles the third-party login process
//format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	if len(segs) < 4 {
		//code will panic if url is malformed
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You must use a valid url. Ex: /auth/login/facebook")
		return
	}
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
			return
		}
		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthUrl for %s: %s", provider, err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
