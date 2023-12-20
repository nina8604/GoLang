package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// Fan out
	xc := fanOut(in, 10)

	// Toruglesshooting
	fmt.Printf("%T \n", xc)
	fmt.Println("*********", len(xc))
	for _, v := range xc {
		fmt.Println("******", <-v)
	}
	for n := range merge(xc...) {
		fmt.Println(n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	var xc []<-chan int // need to be zero
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for n := range in {
			c <- fact(n)
		}
		close(c)
	}()

	return c
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines are
	// done. This must start afte the wg.Add call.

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
