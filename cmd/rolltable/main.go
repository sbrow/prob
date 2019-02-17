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

func main() {
	/*
		n := flag.Int("n", 2, "The number of dice to generate rolls for")
		// d := flag.Int("d", 6, "The number of sides the die has")
		force := flag.Bool("force", false,
			"deletes all old tables before removing new ones")
		flag.Parse()

		if *force {
			if err := dice.DeleteData(); err != nil {
				fmt.Println(err)
			}
		}
		t, err := dice.NewTable(dice.Copy(dice.D6(), *n))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t.Rolls)
	*/
}
