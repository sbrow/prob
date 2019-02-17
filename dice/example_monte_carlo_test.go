package dice

/*
func Example_monteCarlo() {
	// Roll a die 10,000 times.
	n := 10000
	d := Copy(new(D6), 2)
	results := MonteCarlo(d, Average, 3, n)
	fmt.Println(results)

	// Output:
	// 1: 1677
	// 2: 1631
	// 3: 1665
	// 4: 1667
	// 5: 1680
	// 6: 1680
}
*/

/*
func Example_monteCarloWithReroll() {
	rand.Seed(time.Now().UTC().UnixNano())
	// Roll a die 10,000 times.
	results := MonteCarlo(D6(), Average, 10, 1000000)
	max := len(D6().Sides())
	for i := 0; i < max; i++ {
		fmt.Printf("%d: %d\n", i+1, results[i])
	}

	// Output:
	// 1: 1678
	// 2: 1620
	// 3: 1674
	// 4: 1646
	// 5: 1689
	// 6: 1693
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
*/
