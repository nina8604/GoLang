package main

import "fmt"

type person struct {
	first  string
	age    int
	saying string
}

func (p *person) speak() {
	fmt.Println("I'm ", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Nina",
		age:   37,
	}

	p1.speak()
	saySomething(&p1)
	// saySomething(p1) // it doesn't work

}
