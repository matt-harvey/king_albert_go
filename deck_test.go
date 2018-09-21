package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.cards) != 52 {
		t.Fatalf("Expected new deck to contain 52 cards, but contained %d", len(deck.cards))
	}
}

func TestDeckShuffle(t *testing.T) {
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

func TestDeckDeal(t *testing.T) {
	deck := NewDeck()
	numCards := len(deck.cards)
	// to verify that the same card is not dealt twice
	cardsDealt := make(map[Card]bool)
	for i := 0; i != numCards; i++ {
		oldNumCards := len(deck.cards)
		card, dealt := deck.Deal()
		if !dealt {
			t.Fatal("Expected card to be dealt, but it wasn't")
		}
		_, alreadyDealt := cardsDealt[card]
		if alreadyDealt {
			t.Fatalf("Card unexpectedly dealt twice from the same deck: %s", card)
		}
		expectedNumCards := oldNumCards - 1
		actualNumCards := len(deck.cards)
		if actualNumCards != expectedNumCards {
			t.Fatalf(
				"deck.Deal() was expected to change number of cards in deck from %d to %d, but instead changed it to %d",
				oldNumCards,
				expectedNumCards,
				actualNumCards)
		}
		cardsDealt[card] = true
	}
	if len(deck.cards) != 0 {
		t.Fatalf("Expected no cards to be left in deck, but %d cards were left", len(deck.cards))
	}
	card, dealt := deck.Deal()
	expectedCard := Card{Rank(1), Spades}
	if card != expectedCard {
		t.Fatalf("Expected %s, but got %s", expectedCard, card)
	}
	if dealt {
		t.Fatal("Expected card not to be dealt from empty deck, but one was dealt")
	}
}
