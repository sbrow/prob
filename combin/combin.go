// Package combin contains various tools for computing combinatorics.
//
// License
//
// This Program is licensed under the GNU General Public License V3.0.
// A copy of this license is provided with the program, and can be found
// here:
// 	https://github.com/sbrow/combin/blob/master/LICENSE
package combin

import (
	"fmt"
	"sync"
)

type combin struct {
	numer []int
	denom int
}

func Combin(n, r int) *combin {
	return &combin{[]int{n, n - r + 1}, r}
}

func (c *combin) String() string {
	return fmt.Sprintf("(%d...%d) / %d!", c.numer[0], c.numer[1], c.denom)
}

func (c *combin) Float64() float64 {
	return float64(Product(c.numer[1], c.numer[0])) / float64(Fact(c.denom))
}

func (c *combin) Int() int {
	return int(c.Float64())
}

func (c *combin) Div(div combin) float64 {
	return c.Float64() / div.Float64()
}

func (c *combin) Mult(mult ...combin) int {
	prod := c.Int()
	for _, factor := range mult {
		prod *= factor.Int()
	}
	return prod
}

// PermuteR returns all permutations, (with repetition)
// of the characters in the set with length in the interval (min, max).
//
// Output from PermuteR will always match the Regexp "/[\1]+/" where "\1" = set
func PermuteR(set string, min, max int) (combos []string) {
	c := make(chan string)
	var wg sync.WaitGroup
	combos = make([]string, NPRR(len(set), rng(min, max)...))

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

// CombineSets returns the combinations of every rune in left with every rune in right.
//
// Output from CombineSets will always match the Regexp "/[left][right]/"
//
// Each element in combos will always have length 2.
func CombineSets(v ...interface{}) [][]interface{} {
	sets := [][]interface{}{}
	combos := [][]interface{}{}

	for i, val := range v {
		sets = append(sets, []interface{}{})
		switch val.(type) { //reflect.TypeOf(val).Kind() {
		// case reflect.Int:
		case int:
			sets[i] = append(sets[i], val.(int))
		case string:
			for _, r := range val.(string) {
				sets[i] = append(sets[i], string(r))
			}
		case []string:
			for _, r := range val.([]string) {
				sets[i] = append(sets[i], r)
			}
		}
	}

	for _, a := range sets[0] {
		for _, b := range sets[1] {
			combos = append(combos, []interface{}{a, b})
		}
	}
	return combos
}

// Combine returns all combinations of the given interfaces.
func Combine(v ...interface{}) [][2]interface{} {
	combos := [][2]interface{}{}
	for i := 0; i < len(v)-1; i++ {
		for j := i + 1; j < len(v); j++ {
			combos = append(combos, [2]interface{}{v[i], v[j]})
		}
	}
	return combos
}
