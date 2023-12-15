package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(2)
	if y != 14 {
		t.Error("got - ", y, "expended - ", 14)
	}
}

func TestYesrsTwo(t *testing.T) {
	x := YearsTwo(3)
	if x != 21 {
		t.Error("got - ", x, "expected - ", 21)
	}
}

func ExampleYears() {
	n := 2
	y := Years(n)
	fmt.Println(y)
	// Output:
	// 14
}

func ExampleYearsTwo() {
	n := 3
	y := YearsTwo(n)
	fmt.Println(y)
	// Output:
	// 21
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(3)
	}
}
