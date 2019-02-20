package main

import (
	"fmt"
	"strings"
)

const NumColumns = 9
const HandSize = 7

const MinMovementOrigin = 'e'
const MaxMovementOrigin = 't'
const MinMovementDestination = 'a'
const MaxMovementDestination = 'm'

type Board struct {
	Foundations [NumSuits]Foundation `json:"foundations"`
	Success     [NumColumns]Column   `json:"columns"`
	Hand        [HandSize]SpotInHand `json:"hand"`
}

func NewBoard() *Board {
	deck := NewDeck()
	deck.Shuffle()

	board := Board{}

	{
		suitIndex := 0
		WalkSuits(func(suit Suit) {
			board.Foundations[suitIndex] = Foundation{suit, nil}
			suitIndex++
		})
	}

	for i := 0; i != NumColumns; i++ {
		for j := 1; j != i+2; j++ {
			card, dealt := deck.Deal()
			if !dealt {
				panic("Card not dealt when expected during board initialization")
			}
			board.Success[i].Receive(card)
		}
	}

	for i := 0; i != HandSize; i++ {
		card, dealt := deck.Deal()
		if !dealt {
			panic("Card not dealt when expected during board initialization")
		}
		board.Hand[i].Receive(card)
	}

	return &board
}

func (b *Board) LocationAt(label rune) Location {
	switch label {
	case 'a':
		return &b.Foundations[0]
	case 'b':
		return &b.Foundations[1]
	case 'c':
		return &b.Foundations[2]
	case 'd':
		return &b.Foundations[3]
	case 'e':
		return &b.Success[0]
	case 'f':
		return &b.Success[1]
	case 'g':
		return &b.Success[2]
	case 'h':
		return &b.Success[3]
	case 'i':
		return &b.Success[4]
	case 'j':
		return &b.Success[5]
	case 'k':
		return &b.Success[6]
	case 'l':
		return &b.Success[7]
	case 'm':
		return &b.Success[8]
	case 'n':
		return &b.Hand[0]
	case 'o':
		return &b.Hand[1]
	case 'p':
		return &b.Hand[2]
	case 'q':
		return &b.Hand[3]
	case 'r':
		return &b.Hand[4]
	case 's':
		return &b.Hand[5]
	case 't':
		return &b.Hand[6]
	default:
		panic(fmt.Errorf("No board location labelled %c", label))
	}
}

func (b *Board) VictoryState() VictoryState {
	if b.NumLegalMovements() != 0 {
		return Ongoing
	}
	for _, foundation := range b.Foundations {
		card, ok := foundation.ActiveCard()
		if !ok || (card.Rank != MaxRank) {
			return Lost
		}
	}
	return Won
}

func (b *Board) Permits(movement Movement) bool {
	origin := b.LocationAt(movement.Origin)
	destination := b.LocationAt(movement.Destination)
	activeCard, ok := origin.ActiveCard()
	return ok && origin.CanGiveCard() && destination.CanReceive(activeCard)
}

func (b *Board) Execute(movement Movement) {
	if !b.Permits(movement) {
		panic("Illegal move")
	}
	card := b.LocationAt(movement.Origin).GiveCard()
	b.LocationAt(movement.Destination).Receive(card)
}

func (b *Board) maxColumnSize() int {
	max := 0
	for _, column := range b.Success {
		if len(column) > max {
			max = len(column)
		}
	}
	return max
}

func (b *Board) WalkLegalMovements(f func(movement Movement)) {
	for origin := MinMovementOrigin; origin <= MaxMovementOrigin; origin++ {
		for destination := MinMovementDestination; destination <= MaxMovementDestination; destination++ {
			movement := Movement{origin, destination}
			if b.Permits(movement) {
				f(movement)
			}
		}
	}
}

func (b *Board) NumLegalMovements() int {
	result := 0
	b.WalkLegalMovements(func(movement Movement) { result++ })
	return result
}

func (b *Board) String() string {
	fds := b.Foundations
	var bld strings.Builder
	fmt.Fprintf(&bld, "                           a    b    c    d\n")
	fmt.Fprintf(&bld, "____________________________________________\n")
	fmt.Fprintf(&bld, "                          %s  %s  %s  %s\n\n\n", fds[0], fds[1], fds[2], fds[3])
	fmt.Fprintf(&bld, "  e    f    g    h    i    j    k    l    m\n")
	fmt.Fprintf(&bld, "____________________________________________\n")
	for i := 0; i != b.maxColumnSize(); i++ {
		for j, column := range b.Success {
			if j != 0 {
				fmt.Fprintf(&bld, "  ")
			}
			fmt.Fprintf(&bld, "%s", column.PrintableCardAt(i))
		}
		fmt.Fprintf(&bld, "\n")
	}
	fmt.Fprintf(&bld, "\n")
	fmt.Fprintf(&bld, "  n    o    p    q    r    s    t\n")
	fmt.Fprintf(&bld, "____________________________________________\n")
	fmt.Fprintf(&bld, " ")
	for k, spotInHand := range b.Hand {
		if k != 0 {
			fmt.Fprintf(&bld, "  ")
		}
		fmt.Fprintf(&bld, "%s", spotInHand)
	}
	fmt.Fprintf(&bld, "\n\n")

	return bld.String()
}
