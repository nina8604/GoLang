package main

import "fmt"

func main() {
	x := 15
	for {
		fmt.Printf("the x value is %v\n", x)
		if x <= 5 {
			break
		}
		x--
	}
}
