package dice

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Roll returns a random side from each of the given dice.
func Roll(dice ...Die) []string {
	out := make([]string, len(dice))
	for i, d := range dice {
		out[i] = d.Roll()
	}
	return out
}

// Dice is a collection of 0 or more Die structs.
type Dice []Die

// New returns a new set of Dice.
func New(die ...Die) Dice {
	return die
}

// Len returns the number of Dice in the set.
func (d Dice) Len() int {
	return len(d)
}

// Less returns whether the Die at i should be sorted before the Die at j.
func (d Dice) Less(i, j int) bool {
	return len(d[i].Sides) <= len(d[j].Sides)
}

// Name returns a formatted concatenation of the names of each die.
// Names of unlike dice are appended to the name, whereas names of like dice are compressed
// using standard dice notation- i.e. 3, six-sided dice = "3d6".
//
// Names are sorted in descending order, first by quantity of each dice type,
// then by the number of sides each dice type has.
//
// FIXME: Dice with <10 sides are being sorted lower than dice with 10+ sides.
func (d Dice) Name() string {
	names := map[string]int{}
	for _, die := range d {
		names[die.Name]++
	}

	out := []string{}
	for d, n := range names {
		out = append(out, fmt.Sprint(n, d))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(out)))
	return strings.Join(out, Delim)

}

// Roll returns a random side from each Die in the Dice.
func (d Dice) Roll() []string {
	out := make([]string, len(d))
	for i, d := range d {
		out[i] = d.Roll()
	}
	return out
}

// Sort sorts the dice in descending order of number of sides.
func (d Dice) Sort() {
	sort.Sort(sort.Reverse(d))
}

// Swap swaps the elements with indexes i and j.
func (d Dice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// D4 returns a standard, four-sided die.
func D4() Die {
	return Die{Name: "d4", Sides: []string{"1", "2", "3", "4"}}
}

// D6 returns a standard, six-sided die.
func D6() Die {
	return Die{Name: "d6", Sides: []string{"1", "2", "3", "4", "5", "6"}}
}

// D8 returns a standard, eight-sided die.
func D8() Die {
	return Die{Name: "d8", Sides: []string{"1", "2", "3", "4", "5", "6", "7", "8"}}
}

// D10 returns a standard, ten-sided die.
func D10() Die {
	return Die{Name: "d10", Sides: []string{"1", "2", "3", "4", "5", "6", "7", "8",
		"9", "10"}}
}

// D12 returns a standard, twelve-sided die.
func D12() Die {
	return Die{Name: "d12", Sides: []string{"1", "2", "3", "4", "5", "6", "7", "8",
		"9", "10", "11", "12"}}
}

// D20 returns a standard, twelve-sided die.
func D20() Die {
	return Die{Name: "d20", Sides: []string{"1", "2", "3", "4", "5", "6", "7", "8",
		"9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}}
}

// Die represents a physical n-sided die.
type Die struct {
	Name  string
	Sides []string
}

// NewDie returns a new Die type with the given name and sides.
func NewDie(name string, sides ...string) *Die {
	return &Die{Name: name, Sides: sides}
}

// Roll returns a side of the Die at random.
func (d Die) Roll() string {
	return d.Sides[rand.Intn(6)]
}
