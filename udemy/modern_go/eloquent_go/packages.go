package main

import "fmt"

func main() {
	message := "Hello, world!"
	fmt.Printf("%v is of type %T \n", message, message)
	//experiment with format strings and verbs
	s := struct {
		i int
		f float32
	}{i: 3, f: 3.3}
	fmt.Printf("%+v is of type %T \n", s, s)

}

//PRINT takes data and prints it somewhere
//SCAN reads data and parses it into go objects

//when a function starts with F, like Fprint, FprintF, Fscan, this means that the first argument is an io.Writer
//when a funciton ends with F, like Fprintf, Scanf, then that means the second argument is a format string
//when a function begins with an S, like Sprintf, this means that it outputs a string
