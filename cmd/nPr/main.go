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
	"github.com/sbrow/prob/combin"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	n, _ := strconv.Atoi(args[0])
	r := make([]int, len(args[1:]))
	for i, num := range args[1:] {
		r[i], _ = strconv.Atoi(num)
	}
	fmt.Println("nPr: ", combin.NPR(n, r...))
	fmt.Println("nPr^r: ", combin.NPRR(n, r...))
}
