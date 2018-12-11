package hyper_test

import (
	"fmt"

	"github.com/sbrow/prob/combin/hyper"
)

func ExampleHyper_PMF() {
	dist := hyper.Hyper{K: map[string]int{"A": 2, "B": 2}}
	fmt.Printf("%f\n", dist.PMF(map[string]int{"A": 1, "B": 1}))
	// Output:0.666667
}

func ExampleHyper_Sample() {
	deck := map[string]int{
		"Lands":  17,
		"Spells": 23,
	}
	h := hyper.Hyper{K: deck}
	result := h.Sample(map[string]int{"Lands": 3, "Spells": 4})
	fmt.Printf("Dist: map[Lands: %d Spells: %d]\n", result.Dist["Lands"], result.Dist["Spells"])
	fmt.Printf("Sample: map[Lands: %d Spells: %d]\n", result.Sample["Lands"], result.Sample["Spells"])
	fmt.Printf("PMF: %f\n", result.PMF)
	// Output:
	// Dist: map[Lands: 17 Spells: 23]
	// Sample: map[Lands: 3 Spells: 4]
	// PMF: 0.322975
}
