package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.cards) != 52 {
		t.Fatalf("Expected new deck to contain 52 cards, but contained %d", len(deck.cards))
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	oldLen := len(deck.cards)
	deck.Shuffle()
	if len(deck.cards) != oldLen {
		t.Fatalf("Expected shuffled deck to contain %d cards, but contained %d", oldLen, len(deck.cards))
	}
}

func TestDeckWalk(t *testing.T) {
	deck := NewDeck()
	var cards []Card
	deck.Walk(func(card Card) {
		cards = append(cards, card)
	})
	numCards := len(cards)
	if numCards != 52 {
		t.Fatalf("walked %d cards in deck, but expected %d", numCards, 52)
	}
}
