package cards

import (
	"fmt"
)

type PlayingCardDeck struct {
	cards map[PlayingCard]int
}

// New returns a new deck of PlayingCards.
func New() (deck *PlayingCardDeck) {
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	values := []FaceValue{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	cards := make(map[PlayingCard]int)

	for _, suit := range suits {
		for _, value := range values {
			cards[PlayingCard{value, suit}] = 1
		}
	}
	deck.cards = cards
	return deck
}

type MissingCardError struct {
	Card PlayingCard
}

func (m *MissingCardError) Error() string {
	return fmt.Sprintf("card \"%s\" not in deck", m.Card)
}

// Draw removes hands from the deck.
func (p PlayingCardDeck) Draw(hands ...Hander) (*PlayingCardDeck, error) {
	for _, hand := range hands {
		switch v := hand.(type) {
		case PlayingCard:
			if deck, err := p.Remove(v); err != nil {
				return nil, &MissingCardError{v}
			} else {
				return deck, nil
			}
		default:
			// TODO(sbrow): Implement.
		}
	}
	deck := p
	return &deck, nil
}

// Remove removes cards from the deck.
func (p PlayingCardDeck) Remove(cards ...PlayingCard) (*PlayingCardDeck, error) {
	deck := p

	for _, card := range cards {
		if _, ok := deck.cards[card]; !ok {
			return nil, &MissingCardError{card}
		} else {
			delete(deck.cards, card)
		}
	}

	return &deck, nil
}
