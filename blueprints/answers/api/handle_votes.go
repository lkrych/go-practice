package api

import "net/http"

func handleVotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	handleVote(w, r)
}
