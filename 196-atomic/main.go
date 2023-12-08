package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPUs:", runtime.NumCPU())
	fmt.Println("begin Goroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from first func")
		wg.Done()

	}()

	go func() {
		fmt.Println("Hello from second func")
		wg.Done()
	}()

	fmt.Println("mid CPUs:", runtime.NumCPU())
	fmt.Println("mid Goroutines:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("exit of the program")

	fmt.Println("end CPUs:", runtime.NumCPU())
	fmt.Println("end Goroutines:", runtime.NumGoroutine())
}
