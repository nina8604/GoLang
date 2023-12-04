package main

import "fmt"

func main() {
	func() {
		fmt.Println("I create an anonymous func")

	}()
	func(name string, age int) {
		fmt.Println("My name is ", name, "and I'm ", age)
	}("Nina", 37)
}
