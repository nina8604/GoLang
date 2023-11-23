package main

import "fmt"

func main() {
	for y := 0; y < 20; y++ {
		if y%2 == 0 {
			continue
		}
		fmt.Printf("Print only odd number - %v\n", y)
	}
}
