package main

import "fmt"

func main() {
	foo()
	bar("Nina")

	s := aloha("Nina")
	fmt.Println(s)
	s1, n := dogYears(s, 40)
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("I'm from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Miss", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old dog years"), age
}
