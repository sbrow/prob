package combin

import "math"

// NCR returns the number of combinations when choosing r objects from n. rep determines
// whether or not to count with repetition.
//
// If passed more than one r, NCR will return the sum of each combination of n and r.
//
// This means that calling:
// 	NCR(false, 10, 1, 2, 3)
//
// Will return the same result as:
// 	sum := 0
// 	for i := 1; i <= 3; i++ {
// 		sum += NCR(false, n, i)
// 	}
func NCR(rep bool, n int, r ...int) int {
	switch {
	case n <= 0:
		return 0
	case rep:
		return ncrRep(n, r...)
	default:
		return ncr(n, r...)
	}
}

// ncrRep calculates n combinations of r with repetition.
func ncrRep(n int, r ...int) int {
	combos := 0
	for _, k := range r {
		combos += Product(n, n+k-1) / Fact(k)
	}
	return combos
}

// ncr calculates n combinations of r without repetition.
func ncr(n int, r ...int) int {
	combos := 0
	for _, k := range r {
		combos += Product(n-k+1, n) / Fact(k)
	}
	return combos
}

// NPR returns the number of permutations when choosing r objects from n. rep determines
// whether or not to count with repetition.
//
// If passed more than one r, NCR will return the sum of each permutation of n and r.
//
// This means that calling:
// 	NPR(false, 10, 1, 2, 3)
//
// Will return the same result as:
// 	sum := 0
// 	for i := 1; i <= 3; i++ {
// 		sum += NPR(false, n, i)
// 	}
func NPR(rep bool, n int, r ...int) int {
	switch {
	case n <= 0:
		return 0
	case rep:
		return nprRep(n, r...)
	default:
		return npr(n, r...)
	}
}

// npr calculates n permutations of r with repetition.
func nprRep(n int, r ...int) (perm int) {
	for _, num := range r {
		perm += int(math.Pow(float64(n), float64(num)))
	}
	return perm
}

// npr calculates n permutations of r without repetition.
func npr(n int, r ...int) (perm int) {
	for _, k := range r {
		perm += Product(n-k+1, n)
	}
	return perm
}

// Product returns the product of all numbers on the interval [i, n].
func Product(i, n int) (prod int) {
	return ProductFunc(i, n, func(i int, j ...int) int { return i })
}

// ProductFunc returns the product of all numbers on the interval [i, n],
// performing function f on each number.
func ProductFunc(k, n int, f func(int, ...int) int, params ...int) int {
	prod := 1
	for i := k; i <= n; i++ {
		prod *= f(i, params...)
	}
	return prod
}

// SumFunc returns the sum of all numbers on the interval [i, n],
// performing function f on each number.
func SumFunc(k, n int, f func(int, ...int) int, params ...int) int {
	sum := 0
	for i := k; i <= n; i++ {
		sum += f(i, params...)
	}
	return sum
}

// Fact returns the factorial of n. Specifically, it returns Product(1, n).
func Fact(n int) int {
	return Product(1, n)
}
