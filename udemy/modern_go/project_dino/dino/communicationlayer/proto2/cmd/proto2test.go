package main

import (
	"dino/communicationlayer/proto2"
	"flag"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
)

// 1) serialize some data via proto2
// 2) send data via TCP a different service
// 3) deserialize data via proto2 and prtin out the extracted values

// A- TCP client to write data
// B- TCP server to receive the data
func main() {
	//when binary runs, it supports a -op flag EX: ./proto2test -op s => run as server, ./proto2test -op c => run as client
	op := flag.String("op", "s", "s for server, and c for client.")
	switch strings.ToLower(*op) {
	case "s":
		RunProto2Server()
	case "c":
		RunProto2Client()
	}

}

func RunProto2Client() {
	a := &proto2.Animal{
		Id: proto.Int(1),
	}
}

func RunProto2Server() {
	l, err := net.Listen("tcp", ":8282") //create a tcp server
	if err != nil {
		log.Fatal(err)
	}
}
