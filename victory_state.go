package main

type VictoryState int

const (
	Ongoing = VictoryState(iota)
	Lost
	Won
)
