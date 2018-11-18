package poker

import (
	"fmt"
	"testing"
)

var bettests = []struct {
	bets []float64
	odds string
	def  string
}{
	{[]float64{2.0}, "0.5714", "0.4286"},
	{[]float64{2.5}, "0.6250", "0.3750"},
	{[]float64{3.0}, "0.6667", "0.3333"},
	{[]float64{3.5}, "0.7000", "0.3000"},

	{[]float64{3.5, 12}, "0.7059", "0.2941"},
}

var betNtests = []struct {
	bets    []float64
	players int
	out     string
}{
	{[]float64{3.5, 12}, 5, "0.0673"},
}

func TestA(t *testing.T) {
	bets := []float64{3}
	players := 2
	fmt.Println(BetDefN(players, bets...))
	// bets := []float64{3.5, 12, 24}
	// players := 5
	// fmt.Printf("UTG Opens to %.2fbb (%.2f%%)\n", bets[0], 100*BetOdds(bets[0]))
	// fmt.Printf("Pot is %.2fbb\n", SmallBlind+BigBlind+bets[0])
	// fmt.Printf("\n%d Players must defend by 3-Betting %.2f%% (%.2f%%)\n", players,
	// 	100*BetDef(bets[0]), 100*BetDefN(players, bets[0]))
	// fmt.Println(BetOdds(bets[:2]...))
	// fmt.Println(BetDef(bets...))
}

func TestBetOdds(t *testing.T) {
	for i, tt := range bettests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := fmt.Sprintf("%.4f", BetOdds(tt.bets...))
			if got != tt.odds {
				t.Errorf("wanted: %s\ngot: %s", tt.odds, got)
			}
		})
	}
}

func TestBetDef(t *testing.T) {
	for i, tt := range betNtests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := fmt.Sprintf("%.4f", BetDefN(tt.players, tt.bets...))
			if got != tt.out {
				t.Errorf("wanted: %s\ngot: %s", tt.out, got)
			}
		})
	}
}
