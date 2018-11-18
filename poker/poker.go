package poker

import "math"

// BigBlind is the size of the big blind, in big blinds.
var BigBlind = 1.0

// SmallBlind is the size of the small blind, in big blinds.
var SmallBlind = 0.5

// BetDefN returns the percent of the time players must defend against a bet,
// distributed amongst the remaining players.
func BetDefN(players int, bets ...float64) float64 {
	p := BetOdds(bets...)
	return 1 - math.Pow(p, 1.0/float64(players))
}

// BetDef returns the percent of the time players must defend against a bet.
func BetDef(bets ...float64) float64 {
	return 1 - BetOdds(bets...)
}

// BetOdds returns the percent of the time a player can bet.
func BetOdds(bets ...float64) float64 {
	return bets[last(bets)] / sum(append(bets, BigBlind, SmallBlind)...)
}

func last(a []float64) int {
	if len(a) == 1 {
		return 0
	}
	return len(a) - 1
}

func sum(in ...float64) float64 {
	sum := 0.0
	for i := range in {
		sum += in[i]
	}
	return sum
}
