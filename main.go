package main

import (
	"fmt"
)

func main() {
	cards := []string{"Ace of Diamonds", newCard()} // This is a slice
	cards = append(cards, "Six of Spades")          // Returns new slice

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
