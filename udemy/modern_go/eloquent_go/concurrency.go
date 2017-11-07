package main

import (
	"fmt"
	"time"
)

func main() {

	//go routines are a way to execute a piece of code concurrently with other pieces of code

	//this code is executed concurrently to the calling piece of code
	//similar to threads, but they are not real threads
	//they do not correspond to a parallel OS thread
	//goroutines get scheduled and multiplexed amongst open OS threads

	//a channel is a way to communicate between go routines
	// qs := make(chan bool)
	// // go SayHelloFromGoRoutine(qs)
	// go func() { //use anonymous function to do same thing as SayHello...
	// 	fmt.Println("Hello from a third goroutine!")
	// 	qs <- false //captured via closure
	// }()
	// fmt.Println("Hello from main")
	// v := <-qs      //wait on channel and listen for incoming value
	// fmt.Println(v) // prints true

	// select
	ic := make(chan int)
	go periodicSend(ic)
	for i := range ic { //continuously listen to ic channel and intercept val
		fmt.Println(i)
	}
	_, ok := <-ic
	if ok {
		ic <- 3 //will panic because the channel is closed!
	} else {
		fmt.Println("The channel is closed")
	}

}

func SayHelloFromGoRoutine(qs chan bool) {
	fmt.Println("Hello from a new goroutine")
	qs <- true
}

func periodicSend(ic chan int) {
	i := 0
	for i <= 3 {
		ic <- i
		i++
		time.Sleep(1 * time.Second)
	}
	close(ic) //tell goroutine in main to stop listening
}
