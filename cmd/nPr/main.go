/*
Command nPr prints the number of permutations when permuting r objects from n.

imports http://godoc.org/github.com/sbrow/combin

Example:
	$ nPr 10 2
	90
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
	fmt.Println("nPr: ", combin.NPR(false, n, r...))
	fmt.Println("nPr^r: ", combin.NPR(true, n, r...))
}
