package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()} //slice, i.e. Dynamic Array
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
