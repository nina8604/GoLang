package main

import (
	"fmt"
	"math"
)

func main() {
	a := pow(2)
	fmt.Println(a()) // 2
	fmt.Println(a()) // 4
	fmt.Println(a()) // 8
	fmt.Println(a()) // 16
	fmt.Println(a()) // 32

}

func pow(a float64) func() float64 {
	x := 0
	return func() float64 {
		x++
		return math.Pow(a, float64(x))
	}
}
