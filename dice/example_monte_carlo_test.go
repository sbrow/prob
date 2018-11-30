package dice_test

import (
	"fmt"

	"github.com/sbrow/prob/dice"
)

// Run a simulation n times.
func MonteCarlo(n int) map[string]int {
	// The number of times each side was rolled
	counts := map[string]int{}
	die := dice.D6()
	var result string
	for i := 0; i < n; i++ {
		result = die.Roll()
		counts[result]++
	}
	return counts
}

// main()
func Example_monteCarlo() {
	// Roll a die 10,000 times.
	results := MonteCarlo(10000)
	fmt.Println("1:", results["1"])
	fmt.Println("2:", results["2"])
	fmt.Println("3:", results["3"])
	fmt.Println("4:", results["4"])
	fmt.Println("5:", results["5"])
	fmt.Println("6:", results["6"])

	// Output:
	// 1: 1677
	// 2: 1631
	// 3: 1665
	// 4: 1667
	// 5: 1680
	// 6: 1680
}

func Example_monteCarloMany() {
	// Roll a die 10,000 times.
	size := 10
	results := make([]map[string]int, size)
	for i := 0; i < size; i++ {
		results[i] = MonteCarlo(10000)
	}
	fmt.Println(results)
	// Output:
	// 1: 1677
	// 2: 1631
	// 3: 1665
	// 4: 1667
	// 5: 1680
	// 6: 1680
}
