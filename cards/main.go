package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("Length of hand: ", len(hand))
	remainingCards.print()
	fmt.Println("Length of remaining cards: ", len(remainingCards))

}

func newCard() string {
	return "Five of Diamonds"
}
