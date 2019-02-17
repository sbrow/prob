package dice

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Copy returns n copies of Die d.
func Copy(d Die, n int) Dice {
	dice := make([]Die, n)

	for i := 0; i < n; i++ {
		newDie := reflect.New(reflect.ValueOf(d).Elem().Type()).Interface().(Die)
		dice[i] = newDie
	}

	return dice
}

// Dice is a collection of 0 or more Die interfaces.
type Dice []Die

// New returns a new sorted set of Dice.
func New(types ...Die) *Dice {
	// dice := make(Dice, len(types))
	// for i, typ := range types {
	// 	dice[i] = new(typ)
	// }
	dice := Dice(types)
	dice.Sort()
	return &dice
}

// Len returns the number of Dice in the set.
func (d Dice) Len() int {
	return len(d)
}

// Less returns whether the Die at i should be sorted before the Die at j.
func (d Dice) Less(i, j int) bool {
	return len(d[i].Sides()) <= len(d[j].Sides())
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
		names[die.Name()]++
	}

	out := []string{}
	for d, n := range names {
		out = append(out, fmt.Sprint(n, d))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(out)))
	return strings.Join(out, Delim)

}

// Roll randomizes the Up side of each of the dice.
func (d Dice) Roll() []Side {
	sides := make([]Side, len(d))
	for i, die := range d {
		sides[i] = die.Roll()
	}
	return sides
}

// Sort sorts the dice in descending order of number of sides.
func (d *Dice) Sort() {
	sort.Sort(sort.Reverse(d))
}

// Swap swaps the elements with indexes i and j.
func (d Dice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Reroll is a function that takes the result of a dice roll
// and determines which, if any, of the rolled dice should be rerolled.
type Reroll func(Dice) []bool

// Roll returns a random side from each of the dice.
// The dice will be rerolled up to "rerolls" times,
// using the f function to determine which dice to reroll.
func Roll(dice Dice, f Reroll, rerolls int) []Side {
	dice.Roll()

	// Reroll dice until all rerolls are expended,
	// or the roll is satisfactory.
	for i, changed := 0, true; i < rerolls && changed; i++ {
		changed = dice.Reroll(f)
	}
	sides := make([]Side, len(dice))
	for i, die := range dice {
		sides[i] = die.Sides()[*die.Up()]
	}
	return sides
}

func Value(d Die) int {
	if d.Up() == nil {
		return 0
	}
	return d.Sides()[*d.Up()].Value()
}

// Reroll uses fn to determine which dice to reroll, and re-rolls them.
// it returns whether any dice were re-rolled.
func (d Dice) Reroll(fn Reroll) bool {
	if fn == nil {
		return false
	}

	reroll := fn(d)
	changed := false
	for i, die := range d {
		if reroll[i] {
			die.Roll()
			changed = true
		}
	}
	return changed
}

// Sum returns the sum of the rolled dice's Values.
func (d Dice) Sum() int {
	sum := 0
	for _, die := range d {
		sum += Value(die)
	}
	return sum
}

func (d Dice) String() string {
	sides := make([]string, len(d))
	for i, die := range d {
		sides[i] = die.Sides()[*die.Up()].String()
	}
	return fmt.Sprintf("%s=%s", d.Name(), strings.Join(sides, Delim))
}
