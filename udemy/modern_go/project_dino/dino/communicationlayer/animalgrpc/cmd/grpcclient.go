package main

import (
	"context"
	"dino/communicationlayer/animalgrpc"
	"dino/databaselayer"
	"fmt"
	"io"
	"log"

	"flag"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	op := flag.String("op", "s", "s for server, and c for client")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		runGRPCServer()
	case "c":
		runGRPCClient()
	}
}

func runGRPCServer() {
	grpclog.Println("Starting GRPC Server")
	lis, err := net.Listen("tcp", ":8282")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	grpclog.Println("Listening on 127.0.0.1:8282")
	// create new grpc server
	grpcServer := grpc.NewServer()
	// establish server for grpc
	animalServer, err := animalgrpc.NewAnimalGrpcServer(databaselayer.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	//implements grpcserver inteface
	animalgrpc.RegisterAnimalServiceServer(grpcServer, animalServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func runGRPCClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := animalgrpc.NewAnimalServiceClient(conn)
	input := ""
	fmt.Println("All animals? (y/n)")
	fmt.Scanln(&input)
	if strings.ToLower(input) == "y" {
		animals, err := client.GetAllAnimals(context.Background(), &animalgrpc.Request{})
		if err != nil {
			log.Fatal(err)
		}

		for {
			animal, err := animals.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				grpclog.Fatal(err)
			}
			fmt.Println(animal)
		}
		return
	}
	fmt.Println("Nickname?")
	fmt.Scanln(&input)
	a, err := client.GetAnimal(context.Background(), &animalgrpc.Request{Nickname: input})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*a)
}
