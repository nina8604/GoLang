package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Ananymous func ran")
	}()

	func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}("Nina")
}

func foo() {
	fmt.Println("foo ran")
}

//a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) {code}

// an anonymous function
// func(p parameter) (r return(s)) {code}
