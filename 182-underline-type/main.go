package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNums interface {
	constraints.Integer | constraints.Float
}

func addT[T myNums](a, b T) T {
	return a + b
}

type meAlias int

func main() {
	var n meAlias = 12
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))
	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.2))
}
