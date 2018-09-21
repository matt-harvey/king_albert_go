package main

type Location interface {
	CanReceive(Card) bool
	Receive(Card)
	CanGiveCard() bool
	GiveCard() Card
	ActiveCard() (Card, bool)
}
