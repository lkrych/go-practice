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
	jimPointer := &jim
	jim.print()
	jim.updateName("Barry")
	jim.print()
	fmt.Println("Why doesn't this work?")
	// why doesn't this work? see pointers.go
	jimPointer.updateNamePointer("Barry")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
