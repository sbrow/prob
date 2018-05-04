// rolltable enumerates dice probability tables and saves them in csv format.
//
// Usage of rolltable:
//
// -force deletes all old tables before removing new ones.
//
// -n int The number of dice to generate rolls for (default 2).
//
// -d string The name of the die type to roll (default is "d6").
//
// Example:
// 	rolltable -n=3 -d="d6"
// Will generate the rolltable for 3d6.
//
// See https://godoc.org/github.com/sbrow/rolltable for info on the package.
package main

import (
	"flag"
	"fmt"
	"github.com/sbrow/prob/dice"
	"strconv"
)

func main() {
	n := flag.Int("n", 2, "The number of dice to generate rolls for")
	d := flag.Int("d", 6, "The number of sides the die has")
	force := flag.Bool("force", false,
		"deletes all old tables before removing new ones")
	flag.Parse()

	if *force {
		dice.DeleteTables(false)
	}
	dice.Table(*dice.New(fmt.Sprintf("d%d", *d), strconv.Itoa(*d)), *n)
}
