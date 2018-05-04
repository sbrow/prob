package combin

import (
	"fmt"
	"testing"
)

func ExampleNPR() {
	n := 10
	r := []int{1, 2, 3}
	fmt.Println(NPR(n, r...))

	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NPR(n, i)
	}
	fmt.Println(sum)
	//Output:
	//820
	//820
}

func ExampleNPRR() {
	n := 10
	r := []int{1, 2, 3}
	fmt.Println(NPRR(n, r...))

	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NPRR(n, i)
	}
	fmt.Println(sum)
	//Output:
	//1110
	//1110
}

func ExampleSumFunc() {
	// Calculate the 10th fibbonacci number using SumFunc.
	i, n := 0, 10
	fib := func(i int, n ...int) int {
		return NCR(n[0]-i-1, i)
	}
	fmt.Println(SumFunc(i, n, fib, n))
	// Output:
	// 55
}

/*
Tests the speed of Product function implementations.
*/
func BenchmarkProduct(b *testing.B) {
	low, high := 1, 100

	b.ResetTimer()
	b.Run("Single", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Product(low, high)
		}
	})

	b.ResetTimer()
	b.Run("Multi", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			productMult(low, high)
		}
	})
}
