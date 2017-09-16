package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"math/rand"
	"strings"
	"time"
)

// A "deck" is a type, a modified slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{ "Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{ "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	stringSlice := strings.Split(string(bs), ",")
	return deck(stringSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	cardIndices := r.Perm(52)
	
	for i := range d {
		newPosition := cardIndices[i]
		d[i], d[newPosition] = d[newPosition], d[i]	
	}
}