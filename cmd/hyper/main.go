/*
Command hyper lets you create hypergeometric distributions from the commandline.
*/
package main

import (
	"fmt"
	"github.com/sbrow/prob/combin/hyper"
)

const (
	SIZE   = 40
	LANDS  = 17
	SPELLS = SIZE - LANDS
)

func main() {
	deck := map[string]int{
		"k0": LANDS,
		"k1": SPELLS,
	}
	h := hyper.Hyper{K: deck}
	fmt.Println(h.Sample(3, 4))
}
