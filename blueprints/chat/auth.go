package main

import "net/http"

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
