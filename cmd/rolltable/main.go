// rolltable enumerates dice probability tables and saves them in csv format.
//
// Usage of rolltable:
//
// -force deletes all old tables before removing new ones.
//
// Example:
// 	rolltable
// Will generate the rolltable for 2d6.
//
// See package github.com/sbrow/prob/dice for info on the package.
package main

import (
	"flag"

	"github.com/sbrow/prob/dice"
)

func main() {
	// n := flag.Int("n", 2, "The number of dice to generate rolls for")
	// d := flag.Int("d", 6, "The number of sides the die has")
	force := flag.Bool("force", false,
		"deletes all old tables before removing new ones")
	flag.Parse()

	if *force {
		dice.DeleteData()
	}

	dice.NewTable([]dice.Die{dice.D6(), dice.D6()})
}
