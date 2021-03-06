package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vault"
	"vault/pb"

	ratelimitkit "github.com/go-kit/kit/ratelimit"

	"github.com/juju/ratelimit"
	"google.golang.org/grpc"
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

	//listen for termination signals and send an error down errChan
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	rlbucket := ratelimit.NewBucket(1*time.Second, 5)

	//setup endpoints
	hashEndpoint := vault.MakeHashEndpoint(srv)
	{
		hashEndpoint = ratelimitkit.NewTokenBucketThrottler(rlbucket, time.Sleep)(hashEndpoint)
	}
	validateEndpoint := vault.MakeValidateEndpoint(srv)
	{
		validateEndpoint = ratelimitkit.NewTokenBucketThrottler(rlbucket, time.Sleep)(validateEndpoint)
	}
	endpoints := vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}

	//run http server
	go func() {
		log.Println("http:", *httpAddr)
		handler := vault.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	//run gRPC server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler := vault.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterVaultServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()
	//block end of main function and listen until error or system call termination
	log.Fatalln(<-errChan)
}
