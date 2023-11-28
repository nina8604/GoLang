package main

import "fmt"

type person struct {
	first  string
	last   string
	favoIC []string
}

func main() {
	p1 := person{
		first:  "Nina",
		last:   "Iaremenko",
		favoIC: []string{"vanilla", "strawberry"},
	}
	p2 := person{
		first:  "Alexandr",
		last:   "Morozov",
		favoIC: []string{"chocolate", "caramel"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for _, v := range p1.favoIC {
		fmt.Println(p1.first, "favorite is", v)
	}
	for _, v := range p2.favoIC {
		fmt.Println(p2.first, "favorite is", v)
	}
}
