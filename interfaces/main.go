package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//getGreeting involves custom logic so you don't want to use an interface
func (englishBot) getGreeting() string {
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

//the logic is the same here, this would be a good place for the use of interfaces
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
