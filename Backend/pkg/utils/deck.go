package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Spades = iota
	Hearts
	Diamonds
	Clubs
)

// Define card value constants
const (
	Ace = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	suit  int
	value int
}

type Deck []Card

func NewDeck() Deck {
	deck := Deck{}
	for suit := Spades; suit <= Clubs; suit++ {
		for val := Ace; val <= King; val++ {
			deck = append(deck, Card{suit: suit, value: val})
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range *d {
		newPostion := r.Intn(len(*d) - 1)
		(*d)[i], (*d)[newPostion] = (*d)[newPostion], (*d)[i] // swap
	}
}

func (d Deck) ParseDeck() {
	for _, card := range d {
		fmt.Printf("%s of %s\n", cardVal(card.value), cardSuit(card.suit))
	}
}

func cardVal(value int) string {
	switch value {
	case Ace:
		return "Ace"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("%d", value)
	}
}

func cardSuit(suit int) string {
	switch suit {
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	default:
		return "Invalid Suit"
	}
}
