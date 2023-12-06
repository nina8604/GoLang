package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{"Nina"}
	fmt.Println(p1)
	p1 = changeName(p1, "Ninel")
	fmt.Println(p1)

	cangeNamePointer(&p1, "Nino")
	fmt.Println(p1)

}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func cangeNamePointer(p *person, s string) {
	p.first = s
}
