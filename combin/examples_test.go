package combin

import "fmt"

func ExampleSet_Combine_noRepetition() {
	set := NewSet("a", "b", "c")
	fmt.Println(set.Combine(false))

	//Output:
	// a,b
	// a,c
	// b,c
}

func ExampleSet_Combine_repetition() {
	set := NewSet("a", "b", "c")
	fmt.Println(set.Combine(true))

	//Output:
	// a,a
	// a,b
	// a,c
	// b,b
	// b,c
	// c,c
}

func ExampleNCR_repitition() {
	n := 10
	r := []int{1, 2, 3}
	fmt.Println(NCR(true, n, r...))

	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NCR(true, n, i)
	}
	fmt.Println(sum)
	// Output:
	// 285
	// 285
}
func ExampleCombineSets_1() {
	// Make all the combinations of a with c and d,
	// plus all the combinations of b with c and d.
	left, right := NewSet("a", "b"), NewSet("c", "d")
	fmt.Println(CombineSets(*left, *right))
	// Output:
	// a,c
	// a,d
	// b,c
	// b,d
	//
}

func ExampleCombineSets_2() {
	eleven := NewSet("A")
	ten := NewSet("K", "Q", "J", "T").Combine(false)
	hand := CombineSets(*eleven, *ten)
	fmt.Println(hand)

	// Output:
	// A,K,Q
	// A,K,J
	// A,K,T
	// A,Q,J
	// A,Q,T
	// A,J,T
	//
}

func ExampleSet_Join() {
	left, right := NewSet("a", "b"), NewSet("c", "d")
	joined := left.Join(right)
	fmt.Println(joined.Combine(false))

	// Output:
	// a,b
	// a,c
	// a,d
	// b,c
	// b,d
	// c,d
	//
}
func ExamplePermuteR() {
	// Print all permutations of 'a' and 'b'
	// with length between one and two, (using repetition).
	fmt.Println(PermuteR("ab", 1, 2))
	// Output:
	// [a aa ab b ba bb]
}
func ExampleNPR_noRepitition() {
	n := 10
	r := []int{1, 2, 3}
	fmt.Println(NPR(false, n, r...))

	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NPR(false, n, i)
	}
	fmt.Println(sum)
	//Output:
	//820
	//820
}

func ExampleNPR_repitition() {
	n := 10
	r := []int{1, 2, 3}
	fmt.Println(NPR(true, n, r...))

	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NPR(true, n, i)
	}
	fmt.Println(sum)
	//Output:
	//1110
	//1110
}

func ExampleSumFunc() {
	// Calculate the 10th fibbonacci number using SumFunc.
	fib := func(i int, n ...int) int {
		return NCR(false, n[0]-i-1, i)
	}
	n := 10
	fmt.Println(SumFunc(0, n, fib, n))
	// Output:
	// 55
}
