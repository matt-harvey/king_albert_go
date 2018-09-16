package main

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestWalkRanks(t *testing.T) {
	var ranks []Rank
	WalkRanks(func(rank Rank) {
		ranks = append(ranks, rank)
	})
	expected := []Rank{
		Rank(1), Rank(2), Rank(3), Rank(4),
		Rank(5), Rank(6), Rank(7), Rank(8),
		Rank(9), Rank(10), Rank(11), Rank(12),
		Rank(13)}
	if !reflect.DeepEqual(ranks, expected) {
		t.Fatalf("ranks was %v, but expected %v", ranks, expected)
	}
}

func TestWalkSuits(t *testing.T) {
	var suits []Suit
	WalkSuits(func(suit Suit) {
		suits = append(suits, suit)
	})
	expected := []Suit{Spades, Hearts, Diamonds, Clubs}
	if !reflect.DeepEqual(suits, expected) {
		t.Fatalf("suits was %v, but expected %v", suits, expected)
	}
}

func TestSuitString(t *testing.T) {
	// TODO Test panic case
	var str string
	WalkSuits(func(suit Suit) {
		str = suit.String()
		if utf8.RuneCountInString(str) != 1 {
			t.Fatalf("value of String()	for suit %v is %s, which does not contain a exactly 1 rune",
				suit, str)
		}
	})
}

func TestSuitColor(t *testing.T) {
	// TODO Test panic case
	spadesColor := Spades.Color()
	heartsColor := Hearts.Color()
	diamondsColor := Diamonds.Color()
	clubsColor := Clubs.Color()

	if spadesColor != Black {
		t.Fatalf("Spades.Color() is %v, but expected %v", spadesColor, Black)
	}
	if heartsColor != Red {
		t.Fatalf("Hearts.Color() is %v, but expected %v", heartsColor, Red)
	}
	if diamondsColor != Red {
		t.Fatalf("Diamonds.Color() is %v, but expected %v", diamondsColor, Red)
	}
	if clubsColor != Black {
		t.Fatalf("Clubs.Color() is %v, but expected %v", clubsColor, Black)
	}
}

func TestWalkCards(t *testing.T) {
	var cards []Card
	WalkCards(func(card Card) {
		cards = append(cards, card)
	})
	numCards := len(cards)
	if numCards != 52 {
		t.Fatalf("walked %d cards, but expected %d", numCards, 52)
	}
}

func TestCardString(t *testing.T) {
	// All cards should have the same Str width
	WalkCards(func(card Card) {
		cardStr := card.String()
		numRunes := utf8.RuneCountInString(cardStr)
		if numRunes != 3 {
			t.Fatalf("card.String() %s contained %d runes, but was expected to contain 3",
				cardStr, numRunes)
		}
	})
	// Test some samples
	aceSpades := Card{Rank(1), Spades}
	aceSpadesStr := aceSpades.String()
	if aceSpadesStr != " A\u2660" {
		t.Fatalf("Expected String() value of Ace of Spades to be ' A\u2660', but was %s",
			aceSpadesStr)
	}
	queenHearts := Card{Rank(12), Hearts}
	queenHeartsStr := queenHearts.String()
	if queenHeartsStr != " Q\u2661" {
		t.Fatalf("Expected String() value of Queen of Hearts to be ' Q\u2661', but was %s",
			queenHeartsStr)
	}
	tenClubs := Card{Rank(10), Clubs}
	tenClubsStr := tenClubs.String()
	if tenClubsStr != "10\u2663" {
		t.Fatalf("Expected String() value of 10 of Clubs to be '10\u2663', but was %s",
			tenClubsStr)
	}
	eightDiamonds := Card{Rank(8), Diamonds}
	eightDiamondsStr := eightDiamonds.String()
	if eightDiamondsStr != " 8\u2662" {
		t.Fatalf("Expected String() value of 8 of Diamonds to be ' 8\u2662', but was %s",
			eightDiamondsStr)
	}
}
