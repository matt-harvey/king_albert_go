package main

import (
	"math/rand"
	"sync"
	"time"
)

var seedRandOnce sync.Once

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	var deck Deck
	WalkCards(func(card Card) {
		deck.cards = append(deck.cards, card)
	})
	return &deck
}

func (d *Deck) Shuffle() {
	seedRandOnce.Do(func() {
		rand.Seed(time.Now().UTC().UnixNano())
	})
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Deal deals one card from the deck, returning that card and true,
// and reducing the number of cards remaining in the deck by 1; unless
// the deck does not contain any cards, in which case the number of
// cards in the deck remains zero and an undefined card and false are
// returned.
func (d *Deck) Deal() (Card, bool) {
	if len(d.cards) == 0 {
		return Card{}, false
	}
	lim := len(d.cards) - 1
	card := d.cards[lim]
	d.cards = d.cards[:lim]
	return card, true
}
