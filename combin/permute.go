package combin

import "sync"

// PermuteR returns all permutations, (with repetition)
// of the characters in the set with length in the interval (min, max).
//
// Output from PermuteR will always match the Regexp "/[\1]+/" where "\1" = set
func PermuteR(set string, min, max int) (combos []string) {
	c := make(chan string)
	var wg sync.WaitGroup
	combos = make([]string, NPR(true, len(set), rng(min, max)...))

	go func(c chan string) {
		defer wg.Done()
		defer close(c)
		wg.Add(1)
		addRune(c, "", set, min, max)
	}(c)
	wg.Wait()

	i := 0
	for combo := range c {
		combos[i] = combo
		i++
	}
	return combos
}

func rng(low, high int) (arr []int) {
	switch {
	case low == high:
		arr = []int{low}
	case low > high:
		low, high = high, low
		fallthrough
	case low < high:
		arr = make([]int, high-low+1)
		for i := 0; low <= high; low++ {
			arr[i], i = low, i+1
		}
	}
	return arr
}

// addRune is a helper method for PermuteR;
// it recursively generates all permutations of length >= min and <= max,
// from seed and set.
func addRune(c chan string, seed, set string, min, max int) {
	for _, ch := range set {
		switch next := seed + string(ch); {
		case len(next) > max:
			return
		case len(next) >= min:
			c <- next
			fallthrough
		default:
			addRune(c, next, set, min, max)
		}
	}
}
