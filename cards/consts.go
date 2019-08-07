package cards

type FaceValue string

// Enum of valid FaceValues:
const (
	Ace   FaceValue = "A"
	King  FaceValue = "K"
	Queen FaceValue = "Q"
	Jack  FaceValue = "J"
	Ten   FaceValue = "T"
	Nine  FaceValue = "9"
	Eight FaceValue = "8"
	Seven FaceValue = "7"
	Six   FaceValue = "6"
	Five  FaceValue = "5"
	Four  FaceValue = "4"
	Three FaceValue = "3"
	Two   FaceValue = "2"
)

type Suit string

// Enum of valid Suits
const (
	Clubs    Suit = "c"
	Diamonds Suit = "d"
	Hearts   Suit = "h"
	Spades   Suit = "s"
)
