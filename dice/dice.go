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

// A standard, four-sided die.
func D4() Die {
	return Die{Name: "d4", Sides: rng(1, 4)}
}

// A standard, six-sided die.
func D6() Die {
	return Die{Name: "d6", Sides: rng(1, 6)}
}

// A standard, eight-sided die.
func D8() Die {
	return Die{Name: "d8", Sides: rng(1, 8)}
}

// A standard, eight-sided die.
func D10() Die {
	return Die{Name: "d10", Sides: rng(1, 10)}
}

// A standard, twelve-sided die.
func D12() Die {
	return Die{Name: "d12", Sides: rng(1, 12)}
}

// A standard, twenty-sided die.
func D20() Die {
	return Die{Name: "d20", Sides: rng(1, 20)}
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
		arr = []string{strconv.Itoa(low)}
	case low > high:
		low, high = high, low
		fallthrough
	case low < high:
		arr = make([]string, high-low+1)
		for i := 0; low <= high; low++ {
			arr[i], i = strconv.Itoa(low), i+1
		}
	}
	return arr
}
