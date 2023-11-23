package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("The outer loop is %v\t the inner loop is %v\n", i, j)
		}

		fmt.Println()
	}
}
