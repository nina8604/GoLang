package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "\nI'm", p.age)
}

func main() {
	x := person{
		first: "Nina",
		age:   37,
	}
	x.speak()
}
