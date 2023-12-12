package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	const gs = 10

	for j := 0; j < gs; j++ {
		go func() {
			for i := 0; i < gs; i++ {

				c <- i
			}
		}()
	}

	for k := 0; k < gs*gs; k++ {
		fmt.Println(k, <-c)
	}
	fmt.Println("the end")
}
