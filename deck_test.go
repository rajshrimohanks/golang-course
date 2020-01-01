package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckLength := 52

	if len(d) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(d))
	}
}
