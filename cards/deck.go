package main

import "fmt"

// A "deck" is a type, a modified slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}