package main

import (
	"fmt"
)

func main() {

	// c := make(chan int) // doesn't work
	// c <- 42

	// fmt.Println(<-c)

	// buffer
	// c := make(chan int, 1) // buffer

	// c <- 42

	// fmt.Println(<-c)

	// func literal
	// c := make(chan int)

	// go send(c)

	// fmt.Println(<-c)

	// aka
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}

// func send(c chan int) {
// 	c <- 42
// }
