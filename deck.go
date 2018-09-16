package main

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
