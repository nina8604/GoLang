package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println("My name is", d.name, "and I'm walking")
}
func (d *dog) run() {
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
	d1.walk()
	d1.run()
	// youngRun(d1)

	d2 := &dog{"Hans"}
	d2.walk()
	d2.run()
	youngRun(d2)
}
