package dice

import (
	"fmt"
	"strconv"
)

// die represents a physical n-sided die.
type Die struct {
	Name  string
	Sides []string
}

// A standard, six-sided die.
func D6() Die {
	return Die{Name: "d6", Sides: rng(1, 6)}
}

// New returns a new die with the given name and sides.
func New(name string, sides ...string) *Die {
	if len(sides) == 1 {
		fmt.Println(sides[0])
		i, err := strconv.Atoi(sides[0])
		if err != nil {
			panic(err)
		}
		sides = rng(1, i)
	}
	return &Die{Name: name, Sides: sides}
}

func rng(low, high int) (arr []string) {
	switch {
	case low == high:
		arr = []string{string(low)}
	case low > high:
		low, high = high, low
		fallthrough
	case low < high:
		arr = make([]string, high-low+1)
		for i := 0; low <= high; low++ {
			arr[i], i = string(low), i+1
		}
	}
	return arr
}
