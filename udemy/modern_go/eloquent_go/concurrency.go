package main

import "fmt"

func main() {

	//go routines are a way to execute a piece of code concurrently with other pieces of code

	//this code is executed concurrently to the calling piece of code
	//similar to threads, but they are not real threads
	//they do not correspond to a parallel OS thread
	//goroutines get scheduled and multiplexed amongst open OS threads

	//a channel is a way to communicate between go routines
	qs := make(chan bool)
	go SayHelloFromGoRoutine(qs)
	fmt.Println("Hello from main")
	v := <-qs      //wait on channel and listen for incoming value
	fmt.Println(v) // prints true
}

func SayHelloFromGoRoutine(qs chan bool) {
	fmt.Println("Hello from a new goroutine")
	qs <- true
}
