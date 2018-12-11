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
	"reflect"
)

const Delim = ","

type Combo []interface{}

func (c Combo) String() string {
	if len(c) == 1 {
		for _, v := range c {
			return fmt.Sprint(v)
		}
	}
	str := ""
	for i, v := range c {
		str += fmt.Sprintf("%v", v)
		if i+1 != len(c) {
			str += Delim
		}
	}
	return str
}

// Set represents a set of items that can be combined.
type Set []Combo

// NewSet creates a new set from the given items.
func NewSet(v ...interface{}) *Set {
	set := new(Set)
	for _, vv := range v {
		// val := reflect.ValueOf(v)
		// if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		// 	for i := 0; i < val.Len(); i++ {
		// 		set.Add(val.Index(i))
		// 	}
		// } else {
		set.Add(vv)
		// }
	}
	return set
}

// Add ads a new object to the set.
func (s *Set) Add(v interface{}) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		subset := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			subset[i] = val.Index(i)
		}
		*s = append(*s, subset)
	} else {
		*s = append(*s, []interface{}{val})
	}
}

// Combine returns all combinations of the given set with itself.
func (s *Set) Combine(repl bool) *Set {
	if repl {
		return s.combineR()
	}
	combos := []Combo{}
	cpy := []Combo(*s)
	len := len(*s)
	if len == 1 {
		return nil
	}
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			combos = append(combos, []interface{}{cpy[i], cpy[j]})
		}
	}
	set := Set(combos)
	return &set
}

// combineR combines the set with itself (using repetition).
func (s *Set) combineR() *Set {
	combos := []Combo{}
	max := len(*s)
	cpy := []Combo(*s)
	for i := 0; i < max; i++ {
		for j := i; j < max; j++ {
			item := append(cpy[i], cpy[j])
			combos = append(combos, item)
		}
	}
	set := Set(combos)
	return &set
}

// Join adds all the entries from b to the end of s.
func (s *Set) Join(b *Set) *Set {
	temp := append(*s, []Combo(*b)...)
	return &temp
}

// Size returns the number of elements in the set.
func (s *Set) Size() int {
	return len(*s)
}

// CombineSets returns all combinations (without replacement)
// of the items in set a with the items in set b.
func CombineSets(a, b Set) *Set {
	combos := make([]Combo, a.Size()*b.Size())
	i := 0
	for _, aa := range a {
		for _, bb := range b {
			combos[i] = []interface{}{aa, bb}
			i++
		}
	}
	s := Set(combos)
	return &s
}

func (s *Set) String() string {
	str := ""
	for _, v := range *s {
		str += fmt.Sprintln(v)
	}
	return str
}
