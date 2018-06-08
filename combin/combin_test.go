package combin

import (
	"fmt"
	"testing"
)

// Tests the speed of Product function implementations.
func BenchmarkProduct(b *testing.B) {
	low, high := 1, 100

	b.ResetTimer()
	b.Run("Single", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Product(low, high)
		}
	})
}

func TestPermuteR(t *testing.T) {
	want := ("[" +
		"a aa aaa aab aac ab aba abb abc ac aca acb acc " +
		"b ba baa bab bac bb bba bbb bbc bc bca bcb bcc " +
		"c ca caa cab cac cb cba cbb cbc cc cca ccb ccc" +
		"]")
	got := fmt.Sprint(PermuteR("abc", 1, 3))
	if want != got {
		t.Fatalf("wanted: %s\ngot: %s", want, got)
	}
}

func TestCombine(t *testing.T) {
	want := [][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}}
	got := NewSet("a", "b", "c").Combine()
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("wanted: %s\ngot: %s", want, got)
	}
}

var combineSetTests = []struct {
	left  interface{}
	right interface{}
	out   string
}{
	{[]string{"a", "b"}, []string{"c", "d"}, "[[a c] [a d] [b c] [b d]]"},
	// {[]byte{99, 100}, []string{"e", "f"}, "[[c e] [c f] [d e] [d f]]"},
	{[]string{"a", "b"}, 1, "[[a 1] [b 1]]"},
}

func TestCombineSets(t *testing.T) {
	for _, tt := range combineSetTests {
		t.Run(fmt.Sprintf("%s %s", tt.left, tt.right), func(t *testing.T) {
			got := CombineSets(NewSet(tt.left), NewSet(tt.right))
			if fmt.Sprint(got) != tt.out {
				t.Errorf("wanted: %s\ngot: %v", tt.out, got)
			}
		})
	}
}

func TestRNG(t *testing.T) {
	want := []int{3}
	got := rng(3, 3)
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("wanted: %v\ngot: %v\n", want, got)
	}
	want = []int{2, 3, 4}
	got = rng(4, 2)
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Errorf("wanted: %v\ngot: %v\n", want, got)
	}
}

func ncr1(rep bool, n int, r ...int) int {
	if n <= 0 {
		return 0
	}
	combos := 0
	var a, b int

	for _, k := range r {
		if rep {
			a = n
			b = n + k - 1
		} else {
			a = n - k + 1
			b = n
		}
		combos += Product(a, b) / Fact(k)
	}
	return combos
}

func ncr3(rep bool, n int, r ...int) int {
	if n <= 0 {
		return 0
	}

	combos := 0
	if rep {
		for _, k := range r {
			combos += Product(n, n+k-1) / Fact(k)
		}
	} else {
		for _, k := range r {
			combos += Product(n-k+1, n) / Fact(k)
		}
	}
	return combos
}

func ncr4(rep bool, n int, r ...int) int {
	combos := 0
	if n <= 0 {
		return 0
	}
	if rep {
		for _, k := range r {
			combos += Product(n, n+k-1) / Fact(k)
		}
	} else {
		for _, k := range r {
			combos += Product(n-k+1, n) / Fact(k)
		}
	}
	return combos
}