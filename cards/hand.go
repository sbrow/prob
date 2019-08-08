package cards

import "fmt"

// Hand represents a group of cards.
// It is a mapping of card types to quantity (See example).
type Hand map[PlayingCard]int

// Hand ensures that type Hand implements the Hander interface.
func (h Hand) Hand() Hand {
	return h
}

func (h *Hand) Size() int {
	sum := 0
	for _, count := range *h {
		sum += count
	}
	return sum
}

func (h *Hand) String() string {
	str := make([]string, h.Size())
	i := 0
	for card, count := range *h {
		for j := 0; j < count; j++ {
			str[i] = card.String()
			i++
		}
	}
	// return strings.Join(str, ",")
	return fmt.Sprintf("%s", str)
}

// Handers can be represented as a Hand.
type Hander interface {
	Hand() Hand
}
