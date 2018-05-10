package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

//When a server shuts down, you'll want to stop receiving new requests
//save any data to disk, and cleanly end connections with existing open
//connections

//It is recommended that you use an initialization daemon to manage the execution
//of an application

//braintree created the manners package for graceful shutdown, while mainting the same
//interface for listenandserve and the core http package
//it uses WaitGroup from the sync package

func main() {
	handler := newHandler()

	//sets up monitoring of operating system signals
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	fmt.Println("Serving on port 8080")
	manners.ListenAndServe(":8080", handler)
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//handler responding to web requests
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func listenForShutdown(ch <-chan os.Signal) {
	//waits for shutdown signal and reacts
	<-ch
	manners.Close()
}
