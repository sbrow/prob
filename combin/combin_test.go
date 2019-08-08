package combin

import (
	"fmt"
	"testing"
)

func TestPermuteR(t *testing.T) {
	want := "[" +
		"a aa aaa aab aac ab aba abb abc ac aca acb acc " +
		"b ba baa bab bac bb bba bbb bbc bc bca bcb bcc " +
		"c ca caa cab cac cb cba cbb cbc cc cca ccb ccc" +
		"]"
	got := fmt.Sprint(PermuteR("abc", 1, 3))
	if want != got {
		t.Fatalf("wanted: %s\ngot: %s", want, got)
	}
}

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		args *Set
		want string
	}{
		{"a,b,c", NewSet("a", "b", "c"), "a,b\na,c\nb,c\n"},
		{"{a,b}", NewSet([]string{"a", "b"}), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Combine(false)
			switch {
			case got == nil && tt.want != "":
				fallthrough
			case got != nil && got.String() != tt.want:
				t.Errorf("wanted: %s\ngot: %v", tt.want, got)
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

func TestNPR(t *testing.T) {
	type args struct {
		rep bool
		n   int
		r   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"n = -1", args{false, -1, []int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NPR(tt.args.rep, tt.args.n, tt.args.r...); got != tt.want {
				t.Errorf("NPR() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
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
*/

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

func rng2(low, high int) []int {
	var arr []int
	switch {
	case low == high:
		return []int{low}
	case low > high:
		low, high = high, low
		fallthrough
	case low < high:
		for i := low; i <= high; i++ {
			arr = append(arr, i)
		}
	}
	return arr
}

func rng3(low, high int) []int {
	var arr [1]int
	var b []int = arr[:]
	switch {
	case low == high:
		b[0] = low
	case low > high:
		low, high = high, low
		fallthrough
	case low < high:
		// arr = make([]int, high-low+1)
		for i := 0; low <= high; low++ {
			b[i] = low
			i++
		}
	}
	return b
}

func rng4(low, high int) []int {
	var arr []int
	switch {
	case low == high:
		arr = make([]int, 1)
		arr[0] = low
	case low > high:
		low, high = high, low
		fallthrough
	case low < high:
		arr = make([]int, high-low+1)
		for i := 0; low <= high; low++ {
			arr[i] = low
			i++
		}
	}
	return arr
}

func BenchmarkRNG(b *testing.B) {
	var a []int
	m := 0
	b.Run("rng", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			a = rng(0, m)
		}
	})
	b.Run("rng2", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			a = rng2(0, m)
		}
	})
	b.Run("rng3", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			a = rng3(0, m)
		}
	})
	b.Run("rng4", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			a = rng4(0, m)
		}
	})
	fmt.Sprint(a)
}
