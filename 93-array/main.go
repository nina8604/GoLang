package main

import "fmt"

func main() {
	// array literal
	a := [3]int{43, 44, 45}
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers!"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)

	{
		//declare a variable of type [7]int
		var ni [7]int
		// asign a value to index possition zero
		ni[0] = 42
		fmt.Printf("%v \t\t\t %T\n", ni, ni)

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%v \t\t\t %T\n", ni2, ni2)

		// array literal
		// have compoler count elements
		ns := [...]string{"Chocolate", "Vanile", "Strawberry"}
		fmt.Printf("%v \t\t\t %T\n", ns, ns)

		// use bultin len function
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}

}
