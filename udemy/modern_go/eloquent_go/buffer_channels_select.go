package main

import (
	"fmt"
	"time"
)

// buffered channel doesn't block right away when you try to
//send and receive items between the channel. It does this by storing values in a buffer

func main() {
	buffch := make(chan int, 5) //number is buffer size
	buffch <- 3
	buffch <- 2 // channel would lock if I tried to add 4 more items to the channel
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)
	// fmt.Println(<-buffch) //this will also lock	because the channel is empty

	//a select statement is a way to wait for multiple channels either receiving or sending data
	ic := make(chan int)
	select {
	case v1 := <-waitAndSend(3, 2):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 1): //prints 5
		fmt.Println(v2)
	case ic <- 23:
		fmt.Println("ic received a value")
	default: // a way to prevent select statement from blocking code
		fmt.Println("all channels are blocked")
	}
}

func waitAndSend(v, i int) chan int { //will wait for i seconds before sending value v on the return channel
	//this is called a channel generator
	returnCh := make(chan int)
	go func() { //blocking code must be in a new go routine
		time.Sleep(time.Duration(i) * time.Second)
		returnCh <- v
	}()
	return returnCh
}
