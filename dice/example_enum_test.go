package dice_test

import (
	"fmt"

	"github.com/sbrow/prob/dice"
)

func Example_enumeration() {
	d := dice.New(dice.D4(), dice.D4())
	table, err := dice.NewTable(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(table.Rolls)

	// Output:
	// [[1+1] [1+2] [1+3] [1+4] [2+1] [2+2] [2+3] [2+4] [3+1] [3+2] [3+3] [3+4] [4+1] [4+2] [4+3] [4+4]]
}
