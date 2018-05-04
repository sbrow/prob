package combin

import (
	"fmt"
	"testing"
)

func ExampleCombineSets() {
	// Make all the combinations of a with c and d,
	// plus all the combinations of b with c and d.
	left, right := "ab", "cd"
	fmt.Println(CombineSets(left, right))
	// Output:
	// [ac ad bc bd]
}

func ExamplePermuteR() {
	// Print all permutations of 'a' and 'b'
	// with length between one and two, (using repetition).
	fmt.Println(PermuteR("ab", 1, 2))
	// Output:
	// [a aa ab b ba bb]
}

func TestPermuteR(t *testing.T) {
	output := ("[" +
		"a aa aaa aab aac ab aba abb abc ac aca acb acc " +
		"b ba baa bab bac bb bba bbb bbc bc bca bcb bcc " +
		"c ca caa cab cac cb cba cbb cbc cc cca ccb ccc" +
		"]")
	perm := fmt.Sprint(PermuteR("abc", 1, 3))
	if output != perm {
		t.Errorf("\"%s\" does not match.", perm)
		t.Fail()
	}
}
