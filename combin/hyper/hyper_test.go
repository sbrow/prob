package hyper

import (
	"fmt"
	"testing"
)

func ExampleHyper_PMF() {
	dist := Hyper{map[string]int{"A": 2, "B": 2}}
	fmt.Printf("%f\n", dist.PMF([]int{1, 1}))
	// Output:0.666667
}

func TestHyper_PMF(t *testing.T) {
	const OUT = "0.226082"
	dist := &Hyper{map[string]int{"Lands": 17, "Spells": 23}}
	output := fmt.Sprintf("%.6f", dist.PMF([]int{4, 3}))
	if output != OUT {
		fmt.Println("got: ", output)
		fmt.Println("want: ", OUT)
	}
}
