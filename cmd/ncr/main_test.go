package main_test

import (
	"fmt"
	"github.com/sbrow/combin"
)

// Calling:
// 	$ ncr 4 2
func Example() {
	fmt.Println(combin.NCR(4, 2))
	// Output:
	// 6
}
