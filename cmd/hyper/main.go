/*
Command hyper lets you create hypergeometric distributions from the commandline.
*/
package main

import (
	"fmt"

	"github.com/sbrow/prob/combin/hyper"
)

const (
	Size   = 40
	Lands  = 17
	Spells = Size - Lands
)

func main() {
	deck := map[string]int{
		"k0": Lands,
		"k1": Spells,
	}
	h := hyper.Hyper{K: deck}
	fmt.Println(h.Sample(map[string]int{"Lands": 3, "Spells": 4}))
}
