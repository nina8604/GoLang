package main

import "fmt"

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("----------")
	// variatic parametr
	xi = append(xi, 45, 46, 47, 97)

	fmt.Println(xi)
	fmt.Println("----------")
}
