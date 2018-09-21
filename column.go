package main

type Column []Card

func (c *Column) CanReceive(card Card) bool {
	activeCard, ok := c.ActiveCard()
	if !ok {
		return true
	}
	return (activeCard.Color() != card.Color()) && (activeCard.Rank == card.Rank+1)
}

func (c *Column) Receive(card Card) {
	*c = append(*c, card)
}

func (c *Column) CanGiveCard() bool {
	return len(*c) != 0
}

func (c *Column) GiveCard() Card {
	activeCard, ok := c.ActiveCard()
	if !ok {
		panic("Cannot give Card from empty Column")
	}
	*c = (*c)[:len(*c)-1]
	return activeCard
}

func (c *Column) ActiveCard() (Card, bool) {
	numCards := len(*c)
	if numCards == 0 {
		return Card{}, false
	}
	return (*c)[numCards-1], true
}

func (c *Column) PrintableCardAt(i int) string {
	if len(*c) <= i {
		return "   " // TODO bleugh
	}
	ret := (*c)[i].String()
	return ret
}
