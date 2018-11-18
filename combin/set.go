package combin

import "reflect"

// Set represents a set of items that can be combined.
type Set struct {
	elements []interface{}
}

// NewSet creates a new set from the given items.
func NewSet(v ...interface{}) Set {
	set := Set{elements: []interface{}{}}
	for i := range v {
		val := reflect.ValueOf(v[i])
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			for j := 0; j < val.Len(); j++ {
				set.elements = append(set.elements, val.Index(j))
			}
		} else {
			set.elements = append(set.elements, v[i])
		}
	}
	return set
}

// Add ads a new object to the set.
func (s *Set) Add(v interface{}) {
	s.elements = append(s.elements, v)
}

// Combine returns all combinations of the given set with itself.
func (s Set) Combine() [][2]interface{} {
	combos := [][2]interface{}{}
	v := s.elements
	for i := 0; i < len(v)-1; i++ {
		for j := i + 1; j < len(v); j++ {
			combos = append(combos, [2]interface{}{v[i], v[j]})
		}
	}
	return combos
}

// Size returns the number of elements in the set.
func (s *Set) Size() int {
	return len(s.elements)
}

// CombineSets returns all combinations (without replacement)
// of the items in set a with the items in set b.
func CombineSets(a, b Set) [][]interface{} {
	combos := make([][]interface{}, a.Size()*b.Size())
	i := 0
	for _, aa := range a.elements {
		for _, bb := range b.elements {
			combos[i] = []interface{}{aa, bb}
			i++
		}
	}
	return combos
}
