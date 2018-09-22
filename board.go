package main

import (
	"fmt"
	"strings"
)

const numColumns = 9
const handSize = 7

const MinMovementOrigin = 'e'
const MaxMovementOrigin = 't'
const MinMovementDestination = 'a'
const MaxMovementDestination = 'm'

type Board struct {
	foundations [NumSuits]Foundation
	columns     [numColumns]Column
	hand        [handSize]SpotInHand
}

func NewBoard() *Board {
	deck := NewDeck()
	deck.Shuffle()

	board := Board{}

	{
		suitIndex := 0
		WalkSuits(func(suit Suit) {
			board.foundations[suitIndex] = Foundation{suit, nil}
			suitIndex++
		})
	}

	for i := 0; i != numColumns; i++ {
		for j := 1; j != i+2; j++ {
			card, dealt := deck.Deal()
			if !dealt {
				panic("Card not dealt when expected during board initialization")
			}
			board.columns[i].Receive(card)
		}
	}

	for i := 0; i != handSize; i++ {
		card, dealt := deck.Deal()
		if !dealt {
			panic("Card not dealt when expected during board initialization")
		}
		board.hand[i].Receive(card)
	}

	return &board
}

func (b *Board) LocationAt(label rune) Location {
	switch label {
	case 'a':
		return &b.foundations[0]
	case 'b':
		return &b.foundations[1]
	case 'c':
		return &b.foundations[2]
	case 'd':
		return &b.foundations[3]
	case 'e':
		return &b.columns[0]
	case 'f':
		return &b.columns[1]
	case 'g':
		return &b.columns[2]
	case 'h':
		return &b.columns[3]
	case 'i':
		return &b.columns[4]
	case 'j':
		return &b.columns[5]
	case 'k':
		return &b.columns[6]
	case 'l':
		return &b.columns[7]
	case 'm':
		return &b.columns[8]
	case 'n':
		return &b.hand[0]
	case 'o':
		return &b.hand[1]
	case 'p':
		return &b.hand[2]
	case 'q':
		return &b.hand[3]
	case 'r':
		return &b.hand[4]
	case 's':
		return &b.hand[5]
	case 't':
		return &b.hand[6]
	default:
		panic(fmt.Errorf("No board location labelled %c", label))
	}
}

func (b *Board) VictoryState() VictoryState {
	if b.NumLegalMovements() == 0 {
		return Lost
	}
	for _, foundation := range b.foundations {
		card, ok := foundation.ActiveCard()
		if !ok || (card.Rank != MaxRank) {
			return Ongoing
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
	for _, column := range b.columns {
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
	fds := b.foundations
	var bld strings.Builder
	fmt.Fprintf(&bld, "                           a    b    c    d\n")
	fmt.Fprintf(&bld, "____________________________________________\n")
	fmt.Fprintf(&bld, "                          %s  %s  %s  %s\n\n\n", fds[0], fds[1], fds[2], fds[3])
	fmt.Fprintf(&bld, "  e    f    g    h    i    j    k    l    m\n")
	fmt.Fprintf(&bld, "____________________________________________\n")
	for i := 0; i != b.maxColumnSize(); i++ {
		for j, column := range b.columns {
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
	for k, spotInHand := range b.hand {
		if k != 0 {
			fmt.Fprintf(&bld, "  ")
		}
		fmt.Fprintf(&bld, "%s", spotInHand)
	}
	fmt.Fprintf(&bld, "\n\n")

	return bld.String()
}
