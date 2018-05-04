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
	"github.com/sbrow/combin"
	"os"
	// "runtime"
	"strconv"
)

func main() {
	// combin.Tester()
	args := os.Args[1:]
	n, _ := strconv.Atoi(args[0])
	r := make([]int, len(args[1:]))
	for i, num := range args[1:] {
		r[i], _ = strconv.Atoi(num)
	}
	// fmt.Println(combin.PartFactMult(n, r))
	fmt.Println("nPr: ", combin.NPR(n, r...))
	fmt.Println("nPr^r: ", combin.NPRR(n, r...))
}
