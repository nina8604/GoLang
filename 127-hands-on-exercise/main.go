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

	m := make(map[string]person)
	m[p1.last] = p1
	m[p2.last] = p2

	fmt.Println(m)

	for _, v := range m {
		for _, v1 := range v.favoIC {
			fmt.Println(v.first, "favorite is", v1)
		}
	}

}
