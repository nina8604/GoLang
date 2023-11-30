package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("I'm from foo")
}

func bar() {
	fmt.Println("bar")
}
