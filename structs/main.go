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
	var alex person
	var alexContact contactInfo

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact = alexContact
	alex.contact.email = "alex@glad.com"
	alex.contact.zip = 11111

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
