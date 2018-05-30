package combin

import "fmt"

func ExampleCombineSets() {
	// Make all the combinations of a with c and d,
	// plus all the combinations of b with c and d.
	left, right := "ab", "cd"
	fmt.Println(CombineSets(left, right))
	// Output:
	// [[a c] [a d] [b c] [b d]]
}

func ExamplePermuteR() {
	// Print all permutations of 'a' and 'b'
	// with length between one and two, (using repetition).
	fmt.Println(PermuteR("ab", 1, 2))
	// Output:
	// [a aa ab b ba bb]
}
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
