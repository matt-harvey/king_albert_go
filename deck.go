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

func (d *Deck) Walk(f func(card Card)) {
	for _, card := range d.cards {
		f(card)
	}
}
