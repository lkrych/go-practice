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
	// jimPointer := &jim
	// & operator gives memory address of the value this variable is pointing to
	jim.print()
	jim.updateName("Barry")
	jim.print()
	fmt.Println("Why doesn't this work?")
	// why doesn't this work?
	// jimPointer.updateNamePointer("Barry")

	//you can use a shortcut if you don't want to use &,
	//if you define a receiver with a pointer, Golang will convert to pointer
	jim.updateNamePointer("B-MON")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateNamePointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	//the star operator asks for the value at the memory address, this will be the person struct
	//the star in front of the type is a type description
}
