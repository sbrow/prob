/*
Command nCr prints the number of combinations when choosing r objects from n.

Example:
	$ nCr 10 2
	45
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sbrow/prob/combin"
)

func main() {
	args := os.Args[1:]
	n, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	r := make([]int, len(args[1:]))
	for i, num := range args[1:] {
		r[i], err = strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("nCr: ", combin.NCR(false, n, r...))
	fmt.Println("nCr^r: ", combin.NCR(true, n, r...))
}
