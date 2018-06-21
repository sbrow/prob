// Package combin contains various tools for computing combinatorics.
//
// License
//
// This Program is licensed under the GNU General Public License V3.0.
// A copy of this license is provided with the program, and can be found
// here:
// 	https://github.com/sbrow/combin/blob/master/LICENSE
package combin

import "reflect"

// Set represents a set of items that can be combined.
type Set struct {
	Elements []interface{}
}

// NewSet creates a new set from the given items.
func NewSet(v ...interface{}) Set {
	set := Set{Elements: []interface{}{}}
	for i := range v {
		val := reflect.ValueOf(v[i])
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			for j := 0; j < val.Len(); j++ {
				set.Elements = append(set.Elements, val.Index(j))
			}
		} else {
			set.Elements = append(set.Elements, v[i])
		}
	}
	return set
}

// Combine returns all combinations of the given set with itself.
func (s Set) Combine() [][2]interface{} {
	combos := [][2]interface{}{}
	v := s.Elements
	for i := 0; i < len(v)-1; i++ {
		for j := i + 1; j < len(v); j++ {
			combos = append(combos, [2]interface{}{v[i], v[j]})
		}
	}
	return combos
}

// CombineSets returns all combinations (without replacement)
// of the items in set a with the items in set b.
func CombineSets(a, b Set) [][]interface{} {
	combos := [][]interface{}{}
	for _, aa := range a.Elements {
		for _, bb := range b.Elements {
			combos = append(combos, []interface{}{aa, bb})
		}
	}
	return combos
}
