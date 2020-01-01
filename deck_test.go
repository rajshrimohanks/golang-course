package main

import "testing"

func TestNewDeck(t *testing.T) {
	deckLength := 52
	firstCard := "Ace of Spades"
	lastCard := "King of Clubs"

	d := newDeck()

	if len(d) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected first card of '%v', but got %v", firstCard, d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected last card of '%v', but got %v", lastCard, d[len(d)-1])
	}
}
