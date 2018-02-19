package vaultd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)

	flag.Parse()
	ctx := context.Background()
	srv := vault.NewService()
	errChan := make(chan error)

	//here in a new go routine, we ask signal.Notify to tell us when we receive a SIGINT
	//when this happens, this signal will be sent down the c channel
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	hashEndpoint := vault.MakeHashEndpoint(srv)
	validateEndpoint := vault.MakeValidateEndpoint(srv)

	//endpoints is a wrapper around our service
	endpoints := vault.Endpoints{
		HashEndpoint: hashEndpoint,
		ValidateEndpoint, validateEndpoint,
	}

	//HTTP transport
	go func() {
		log.Println("http:", *httpAddr)
		handler := vault.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

}
