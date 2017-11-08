package server

import (
	"fmt"
	"net/http"
)

//RunServer sets up an HTTP server on whatever address is passed to it
func RunServer(addr string) error {
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(addr, nil) //second argument says that we will use default multiplexor

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the server %s!", r.RemoteAddr)
}
