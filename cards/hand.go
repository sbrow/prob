package cards

// Hand represents a group of cards.
// It is a mapping of card types to quantity (See example).
type Hand map[string]int

// Hand ensures that type Hand implements the Hander interface.
func (h Hand) Hand() Hand {
	return h
}

// Handers can be represented as a Hand.
type Hander interface {
	Hand() Hand
}
