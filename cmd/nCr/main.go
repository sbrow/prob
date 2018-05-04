/*
Command nCr prints the number of combinations when choosing r objects from n.

imports http://godoc.org/github.com/sbrow/combin

Example:
	$ nCr 10 2
	45
*/
package main

import (
	"fmt"
	"github.com/sbrow/combin"
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
	fmt.Println("nCr: ", combin.NCR(n, r...))
	fmt.Println("nCr^r: ", combin.NCRr(n, r...))
}
