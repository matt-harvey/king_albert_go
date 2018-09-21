package main

type SpotInHand struct {
	card    Card
	hasCard bool
}

func (s SpotInHand) String() string {
	if s.hasCard {
		return s.card.String()
	}
	return "   "
}

func (s *SpotInHand) CanReceive(card Card) bool {
	return false
}

func (s *SpotInHand) Receive(card Card) {
	s.card = card
	s.hasCard = true
}

func (s *SpotInHand) CanGiveCard() bool {
	return s.hasCard
}

func (s *SpotInHand) GiveCard() Card {
	if !s.hasCard {
		panic("SpotInHand does not contain a card")
	}
	s.hasCard = false
	return s.card
}

func (s *SpotInHand) ActiveCard() (Card, bool) {
	if s.hasCard {
		return s.card, true
	}
	return Card{}, false
}
