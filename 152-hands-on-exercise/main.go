package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := foo(x...)
	fmt.Println(r)
	fmt.Println(bar(x))
}

func foo(ii ...int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func bar(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
