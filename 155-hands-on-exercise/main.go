package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	weidth float64
}

type cirle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.weidth
}

func (c cirle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2) * 2
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println("the area - ", sh.area())
}

func main() {
	s1 := square{
		length: 4,
		weidth: 4,
	}

	c1 := cirle{
		radius: 3,
	}

	info(s1)
	info(c1)
}
