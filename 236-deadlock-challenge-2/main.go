package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// c <- 1

	// my solution
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)

	}

	// fmt.Println(<-c)
}
