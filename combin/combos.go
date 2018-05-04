package combin

import (
	"github.com/sbrow/ranges"
	"sync"
)

/*
PermuteR returns all permutions, (with repetition) of the characters in set
with length between min and max (inclusive).

Output from PermuteR will always match the Regexp "/[set]+/"
*/
func PermuteR(set string, min, max int) (combos []string) {
	c := make(chan string)
	var wg sync.WaitGroup
	combos = make([]string, NPRR(len(set), ranges.Enum(min, max)...))

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

/*
TODO: Implement, then make public.
*/
func permute(set string, min, max int) (combos []string) {
	// c := make(chan string)
	return combos
}

/*
addItem is a helper method for PermuteR;
it recursively generates all permutations of length >= min and <= max,
from seed and set.
*/
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

/*
CombineSets returns the combinations of every rune in left with every rune in right.

Output from CombineSets will always match the Regexp "/[left][right]/"

Each element in combos will always have length 2.
*/
func CombineSets(left, right string) (combos []string) {
	combos = make([]string, len(left)*len(right))
	i := 0
	for _, l := range left {
		for _, r := range right {
			combos[i] = string(l) + string(r)
			i++
		}
	}
	return combos
}

// Deprecated: Not particularly useful, but has not yet been thrown out.
func combine(curr string, next []string, length ...int) (combos []string) {
	switch {
	case len(curr) == 0:
		combos = next
	case len(length) > 0 && len(curr) == length[0]:
		combos = []string{curr}
	default:
		combos = make([]string, len(next))
		for i := 0; i < len(next); i++ {
			combos[i] = curr + next[i]
		}
	}
	return combos
}
