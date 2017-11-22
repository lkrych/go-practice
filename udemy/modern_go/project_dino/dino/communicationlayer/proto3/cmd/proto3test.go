package main

import (
	"dino/communicationlayer/proto3"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/golang/protobuf/proto"
)

// 1) serialize some data via proto3
// 2) send data via TCP a different service
// 3) deserialize data via proto3 and prtin out the extracted values

// A- TCP client to write data
// B- TCP server to receive the data
func main() {
	//when binary runs, it supports a -op flag EX: ./proto3test -op s => run as server, ./proto3test -op c => run as client
	op := flag.String("op", "s", "s for server, and c for client ")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto3Server()
	case "c":
		RunProto3Client()
	}

}

//RunProto3Client send Animal protocol buffer to localhost 8282
func RunProto3Client() {
	a := &proto3.Animal{
		Id:         1,
		AnimalType: "Butterfly",
		Nickname:   "Ludwig",
		Zone:       3,
		Age:        20,
	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	SendData(data)
}

//SendData create a TCP server and send data to localhost 8282
func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8282")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)
}

//RunProto3Server Listen for TCP messages on localhost 8282
func RunProto3Server() {
	l, err := net.Listen("tcp", ":8282") //create a tcp server
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept() //wait until receipt of tcp request, provides a connection object for request
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(cn net.Conn) {
			defer cn.Close()
			data, err := ioutil.ReadAll(cn)
			if err != nil {
				return
			}
			a := &proto3.Animal{}
			err = proto.Unmarshal(data, a)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(a)
		}(c)
	}
}
