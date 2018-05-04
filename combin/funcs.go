package combin

import (
	"math"
	"sync"
)

/*
NCR returns the number of combinations when choosing r objects from n
without replacement.

If passed more than one r, NCR will return the sum of each nPr permutation.

This means that calling:
	NCR(10, 1, 2, 3)

Will return the same result as:
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NCR(n, i)
	}
*/
func NCR(n int, r ...int) (combin int) {
	if n <= 0 {
		return 0
	}
	for _, k := range r {
		combin += Product(n-k+1, n) / Fact(k)
	}
	return combin
}

/*
NCRR returns the number of combinations when choosing r objects from n
withreplacement.

If passed more than one r, NCRR will return the sum of each nPr permutation.

This means that calling:
	NCRR(10, 1, 2, 3)

Will return the same result as:
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NCRR(n, i)
	}
*/
func NCRR(n int, r ...int) (combin int) {
	if n <= 0 {
		return 0
	}
	for _, k := range r {
		combin += Product(n, n+k-1) / Fact(k)
	}
	return combin
}

/*
NPR returns the number of permutations when permuting r objects from n
without replacement.

If passed more than one r, NPR will return the sum of each nPr permutation.

This means that calling:
	NPR(10, 1, 2, 3)

Will return the same result as:
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NPR(n, i)
	}
*/
func NPR(n int, r ...int) (perm int) {
	if n < 0 {
		return 0
	}
	for _, k := range r {
		perm += Product(n-k+1, n)
	}
	return perm
}

/*
NPRR returns the number of permutations when permuting r objects from n
with replacement.

If passed more than one r, NPRR will return the sum of each nPr permutation.

This means that calling:
	NPRR(10, 1, 2, 3)

Will return the same result as:
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NPRR(n, i)
	}
*/
func NPRR(n int, r ...int) (perm int) {
	if n <= 0 {
		return 0
	}
	for _, num := range r {
		perm += int(math.Pow(float64(n), float64(num)))
	}
	return perm
}

//Product returns the product of all numbers between i and n (inclusive).
func Product(i, n int) (prod int) {
	return ProductFunc(i, n, func(i int, j ...int) int { return i })
}

/*
ProductFunc returns the product of all numbers between i and n (inclusive).
Performing function f on each number.
*/
func ProductFunc(k, n int, f func(int, ...int) int, params ...int) (prod int) {
	prod = 1
	for i := k; i <= n; i++ {
		prod *= f(i, params...)
	}
	return prod
}

/*
SumFunc returns the sum of all numbers between k and n (inclusive). Performing
function f on each number.
*/
func SumFunc(k, n int, f func(int, ...int) int, params ...int) (sum int) {
	for i := k; i <= n; i++ {
		sum += f(i, params...)
	}
	return sum
}

/*
PartFactMult calculates the product of all numbers between n and lim using two
goroutines.

Deprecated: Slow
*/
func productMult(n, lim int) (pfact int) {
	var wg sync.WaitGroup
	wg.Add(2)
	m := (n - lim) / 2
	a := n
	b := m

	go func() {
		defer wg.Done()
		for i := n - 1; i > m; i-- {
			a *= i
		}
	}()
	go func() {
		defer wg.Done()
		for j := m; j >= lim; j-- {
			b *= m
		}
	}()
	wg.Wait()

	return a * b
}

/*
Fact returns the factorial of n.

Equivilant to calling Product(1, n)
*/
func Fact(n int) (nFact int) {
	return Product(1, n)
}
