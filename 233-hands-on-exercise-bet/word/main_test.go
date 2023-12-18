package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	s := UseCount("hello hello world")
	for k, v := range s {
		switch k {
		case "hello":
			if v != 2 {
				t.Error("Got - ", v, "expexted - ", 2)
			}
		case "world":
			if v != 1 {
				t.Error("Got - ", v, "expexted - ", 1)
			}
		}
	}

}

func TestCount(t *testing.T) {
	s := "hello hello world"
	xs := Count(s)
	if xs != 3 {
		t.Error("Got - ", xs, "expexted - ", 3)
	}
}

func ExampleCount() {
	s := "hello hello world"
	xs := Count(s)
	fmt.Println(xs)
	// Output:
	// 3
}

func BenchmarkUseCount(b *testing.B) {
	s := "hello hello world"
	for i := 0; i < b.N; i++ {
		UseCount(s)
	}
}

func BenchmarkCount(b *testing.B) {
	s := "hello hello world"
	for i := 0; i < b.N; i++ {
		Count(s)
	}
}
