package cards

import "fmt"

// PlayingCard represents a playing card from a typical 52-card deck of cards.
type PlayingCard struct {
	FaceValue
	Suit
}

// Hand returns the Card as a Hand interface.
func (p PlayingCard) Hand() (hand Hand) {
	hand = make(map[PlayingCard]int)
	hand[p] = 1
	return hand
}

func (p PlayingCard) String() string {
	return fmt.Sprintf("%s%s", p.FaceValue, p.Suit)
}
