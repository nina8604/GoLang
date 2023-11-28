package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk:   true,
		first: "ETHAN HANK",
	}
	p2 := person{
		first: "Janny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T\t %#v\n", sa1, sa1)
	fmt.Printf("%T\t %#v\n", p2, p2)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(sa1.first, sa1.person.first)

}
