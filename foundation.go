package main

import "fmt"

type Foundation struct {
	Suit  Suit   `json:"suit"`
	Cards []Card `json:"cards"`
}

func (f Foundation) String() string {
	activeCard, ok := f.ActiveCard()
	if !ok {
		return fmt.Sprintf("  %s", f.Suit)
	}
	return activeCard.String()
}

func (f *Foundation) CanReceive(card Card) bool {
	return (card.Suit == f.Suit) && (card.Rank == f.nextRank())
}

func (f *Foundation) Receive(card Card) {
	f.Cards = append(f.Cards, card)
}

func (f *Foundation) CanGiveCard() bool {
	return false
}

func (f *Foundation) GiveCard() Card {
	panic("Cannot give Card from Foundation")
}

func (f *Foundation) ActiveCard() (Card, bool) {
	if len(f.Cards) == 0 {
		return Card{}, false
	}
	return f.Cards[len(f.Cards)-1], true
}

func (f *Foundation) nextRank() Rank {
	activeCard, ok := f.ActiveCard()
	if !ok {
		return Rank(1)
	}
	return activeCard.Rank + 1
}
