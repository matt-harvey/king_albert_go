package main

type SpotInHand struct {
	Card    Card `json:"card"`
	HasCard bool `json:"has_card"`
}

func (s SpotInHand) String() string {
	if s.HasCard {
		return s.Card.String()
	}
	return "   "
}

func (s *SpotInHand) CanReceive(card Card) bool {
	return false
}

func (s *SpotInHand) Receive(card Card) {
	s.Card = card
	s.HasCard = true
}

func (s *SpotInHand) CanGiveCard() bool {
	return s.HasCard
}

func (s *SpotInHand) GiveCard() Card {
	if !s.HasCard {
		panic("SpotInHand does not contain a card")
	}
	s.HasCard = false
	return s.Card
}

func (s *SpotInHand) ActiveCard() (Card, bool) {
	if s.HasCard {
		return s.Card, true
	}
	return Card{}, false
}
