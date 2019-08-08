package cards

import (
	"fmt"
)

type PlayingCardDeck struct {
	cards map[PlayingCard]int
	odds  *DeckOdds
}

// New returns a new deck of PlayingCards.
func New() (deck *PlayingCardDeck) {
	deck = &PlayingCardDeck{cards: make(map[PlayingCard]int)}
	deck.odds = newDeckOdds(deck)
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	values := []FaceValue{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	for _, suit := range suits {
		for _, value := range values {
			deck.cards[PlayingCard{value, suit}] = 1
		}
	}
	return deck
}

type MissingCardError struct {
	Card PlayingCard
}

func (m *MissingCardError) Error() string {
	return fmt.Sprintf("card \"%s\" not in deck", m.Card)
}

// Draw removes hands from the deck.
func (p PlayingCardDeck) Draw(hands ...Hander) (deck *PlayingCardDeck, err error) {
	deck, _ = p.Remove()
	for _, hand := range hands {
		switch v := hand.(type) {
		case PlayingCard:
			if deck, err = p.Remove(v); err != nil {
				return nil, &MissingCardError{v}
			} else {
				return deck, nil
			}
		case Hand:
			for card, count := range v {
				cards := make([]PlayingCard, count)
				for i, _ := range cards {
					cards[i] = card
				}
				if deck, err = deck.Remove(cards...); err != nil {
					return nil, &MissingCardError{card}
				}
			}
		default:
			panic(v)
		}
	}
	return deck, nil
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

func (p *PlayingCardDeck) Size() int {
	sum := 0
	for _, copies := range p.cards {
		sum += copies
	}
	return sum
}
