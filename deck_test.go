package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.cards) != 52 {
		t.Fatalf("Expected new deck to contain 52 cards, but contained %d", len(deck.cards))
	}
}
