package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zip:   94000,
		},
	}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
