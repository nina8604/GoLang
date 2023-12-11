package main

import "fmt"

func main() {
	// c := make(chan<- int, 2) // it doesn't work
	// c := make(<-chan int, 2) // it doesn't work too

	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	// c <- 42
	// c <- 43

	// fmt.Println(<-c)
	// fmt.Println(<-c)

	fmt.Println("------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// general to specific convert
	fmt.Println("------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// specific to general doesn't convert
	fmt.Println("------")
	fmt.Printf("c\t%T\n", (chan int)(cr))
	fmt.Printf("c\t%T\n", (chan int)(cs))

}
