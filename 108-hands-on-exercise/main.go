package main

import "fmt"

func main() {
	x := make([]int, 10)
	for i := 0; i < 10; i++ {
		x[i] = 42 + i
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}

	fmt.Printf("%T \t %#v\n", x, x)
	fmt.Printf("%T \t %v\n", x, x)

	fmt.Println("---------")

	fmt.Printf("x - %#v\n", x[:5])
	fmt.Println("---------")

	fmt.Printf("x - %#v\n", x[5:])
	fmt.Println("---------")

	fmt.Printf("x - %#v\n", x[2:7])
	fmt.Println("---------")

	fmt.Printf("x - %#v\n", x[1:6])
	fmt.Println("---------")
}
