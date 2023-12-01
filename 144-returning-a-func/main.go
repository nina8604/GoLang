package main

import "fmt"

func main() {
	fmt.Println(foo())

	y := bar()
	fmt.Println(y)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

func foo() int {
	return 42
}
func bar() int {
	return func() int {
		return 43
	}()
}
