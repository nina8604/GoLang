package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// Fan out
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	// Fan in
	var y int
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 30; j < 40; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
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
