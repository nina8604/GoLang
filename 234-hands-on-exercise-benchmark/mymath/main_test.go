package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{1, 4, 6, 8, 100}, 6.0},
		test{[]int{0, 8, 10, 1000}, 9.0},
		test{[]int{9000, 4, 10, 8, 6, 12}, 9.0},
		test{[]int{123, 744, 140, 200}, 170.0},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected -", v.answer, "Got -", x)
		}
	}

}

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}
	fmt.Println(CenteredAvg(a))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	a := []int{1, 4, 6, 8, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(a)
	}
}
