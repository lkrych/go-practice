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
