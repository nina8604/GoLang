package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x)
	fmt.Println(y, z)
}

func foo() int {
	return 37
}

func bar() (int, string) {
	return 32, fmt.Sprint("Alex")
}
