package cards

import "github.com/sbrow/prob/combin"

type DeckOdds struct {
	deck    *PlayingCardDeck
	success int
	total   int
}

func NewDeckOdds(deck *PlayingCardDeck) *DeckOdds {
	odds := &DeckOdds{
		deck: deck,
	}
	odds.deck = deck
	odds.deck.odds = odds

	return odds
}

// Q: what should the following return?
// 	dec.odds.Draw("Ah")
// A:
//  - The odds of drawing "Ah"
//  - A copy of the deck minus "Ah"
// TODO(sbrow): Convert card into ...Hands
func (d *DeckOdds) Draw(hands ...Hander) (odds DeckOdds, err error) {
	newDeck := *d.deck
	odds = *NewDeckOdds(&newDeck)

	success, total := 1, 1
	for _, h := range hands {
		hand := h.Hand()
		// get number of successful combos for each hand.
		for card, _ := range hand {
			cards, ok := odds.deck.cards[card]
			if !ok {
				return *d, &MissingCardError{card}
			}
			success *= cards
		}
		total *= combin.NCR(false, odds.deck.Size(), hand.Size())
		if odds.deck, err = odds.deck.Draw(hand); err != nil {
			return *d, err
		}
	}
	odds.success = success
	odds.total = total

	return odds, nil
}

func (d *DeckOdds) Float64() float64 {
	return float64(d.success) / float64(d.total)
}
