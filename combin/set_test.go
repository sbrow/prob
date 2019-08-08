package combin

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want string
	}{
		{"{{a,b}}", []interface{}{[]string{"a", "b"}}, "a,b\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(tt.args...); got.String() != tt.want {
				t.Errorf("NewSet() = \"%v\", want %v", got, tt.want)
			}
		})
	}
}

func TestCombineSets(t *testing.T) {

	combineSetTests := []struct {
		sets []Set
		out  string
	}{
		{[]Set{*NewSet("a", "b"), *NewSet("c", "d")}, "a,c\na,d\nb,c\nb,d\n"},
		{[]Set{*NewSet([]string{"a", "b"}), *NewSet("c")}, "a,b,c\n"},
		{[]Set{*NewSet([]byte{99, 100}), *NewSet("e", "f")}, "99,100,e\n99,100,f\n"},
		{[]Set{*NewSet("a", "b"), *NewSet(1)}, "a,1\nb,1\n"},
		{[]Set{*NewSet("a", "b"), *NewSet(1), *NewSet(1, 2)}, strings.Join([]string{"a,1,1", "a,1,2", "b,1,1", "b,1,2"}, "\n") + "\n"},
	}
	for _, tt := range combineSetTests {
		t.Run("", func(t *testing.T) {
			got := CombineSets(tt.sets...)
			if fmt.Sprint(got) != tt.out {
				t.Errorf("wanted: %s\ngot: %v", tt.out, got)
			}
		})
	}
}
