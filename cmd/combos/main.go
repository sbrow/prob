// Command combos returns the combinations of all elements in a given set.
package main

import (
	"fmt"
	"github.com/sbrow/prob/combin"
)

func main() {
	fmt.Println(combin.PermuteR("abcde", 2, 3))
}
