package dice_test

import (
	"fmt"
	"strconv"

	"github.com/sbrow/prob/dice/v2"
)

// Run a simulation n times.
func MonteCarlo(n int) [6]int {
	// The number of times each side was rolled
	count := [6]int{0, 0, 0, 0, 0, 0}
	die := dice.D6()
	var result string
	for i := 0; i < n; i++ {
		result = die.Roll()
		intR, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		count[intR-1]++
	}
	return count
}

// main()
func Example_monteCarlo() {
	// Roll a die 10,000 times.
	fmt.Println(MonteCarlo(10000))

	// Output:
	// [1678 1631 1666 1667 1680 1678]
}
