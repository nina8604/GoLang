package main

import (
	"fmt"
	"mymodule/226-hands-on-exercise-documentation/doggy"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  doggy.Years(10),
	}
	fmt.Println(fido)
}
