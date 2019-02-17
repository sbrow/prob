package dice

import (
	"strings"
)

// Run a simulation n times.
func MonteCarlo(dice Dice, fn Reroll, rerolls, n int) map[string]int {
	// The number of times each side was rolled
	counts := map[string]int{}

	for i := 0; i < n; i++ {
		roll := Roll(dice, fn, rerolls)
		str := make([]string, len(roll))
		for j, side := range roll {
			str[j] = side.String()
		}
		counts[strings.Join(str, Delim)]++
	}
	return counts
}
