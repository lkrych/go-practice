package main

import "fmt"

// A "deck" is a type, a modified slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{ "Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{ "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}