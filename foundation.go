package main

import "fmt"

type Foundation struct {
	suit  Suit
	cards []Card
}

func (f Foundation) String() string {
	activeCard, ok := f.ActiveCard()
	if !ok {
		return fmt.Sprintf("  %s", f.suit)
	}
	return activeCard.String()
}

func (f *Foundation) CanReceive(card Card) bool {
	return (card.Suit == f.suit) && (card.Rank == f.nextRank())
}

func (f *Foundation) Receive(card Card) {
	f.cards = append(f.cards, card)
}

func (f *Foundation) CanGiveCard() bool {
	return false
}

func (f *Foundation) GiveCard() Card {
	panic("Cannot give Card from Foundation")
}

func (f *Foundation) ActiveCard() (Card, bool) {
	if len(f.cards) == 0 {
		return Card{}, false
	}
	return f.cards[len(f.cards)-1], true
}

func (f *Foundation) nextRank() Rank {
	activeCard, ok := f.ActiveCard()
	if !ok {
		return Rank(1)
	}
	return activeCard.Rank + 1
}
