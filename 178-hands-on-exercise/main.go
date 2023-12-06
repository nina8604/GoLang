package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println("My name is", d.name, "and I'm walking")
}
func (d dog) run() {
	d.name = "Rover"
	fmt.Println("My name is", d.name, "and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}
func youngWalk(y youngin) {
	y.walk()
}

func main() {
	d1 := dog{"Sema"}
	youngRun(d1)

	d2 := dog{"Hans"}
	youngRun(d2)
	d2 = d2.changeName("Harry")
	youngRun(d2)

}

func (d dog) changeName(s string) dog {
	d.name = s
	return d
}
