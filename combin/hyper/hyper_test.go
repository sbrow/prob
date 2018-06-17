package hyper

import (
	"fmt"
	"testing"
)

var sampletests = []struct {
	sample map[string]int
	out    string
}{
	{map[string]int{"A": 2, "B": 0}, "0.16667"},
	{map[string]int{"A": 0, "B": 2}, "0.16667"},
	{map[string]int{"A": 1, "B": 1}, "0.66667"},
}

func TestHyperSample(t *testing.T) {
	dist := Hyper{map[string]int{"A": 2, "B": 2}}
	for _, tt := range sampletests {
		t.Run(fmt.Sprint(tt.sample), func(t *testing.T) {
			got := dist.Sample(tt.sample).PMF
			if fmt.Sprintf("%.5f", got) != tt.out {
				t.Errorf("wanted: %s\ngot: %.5f\n", tt.out, got)
			}
		})
	}
}
