package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println(PrintSquare(square, x))
}

func square(n int) int {
	return n * n
}

func PrintSquare(f func(int) int, n int) string {
	x := f(n)
	return fmt.Sprintf("the number %v, aquare is %v", n, x)
}
