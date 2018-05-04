// Command combos returns the combinations of all elements in a given set.
package main

import (
	"fmt"
	"github.com/sbrow/combin"
)

func main() {
	fmt.Println(combin.Substrings("abcde", 2, 3))
}
