package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()} // This is a slice
	cards = append(cards, "Six of Spades")          // Returns new slice

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
