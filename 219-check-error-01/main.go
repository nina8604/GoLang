package main

import (
	"fmt"
)

func main() {

	n, err := fmt.Println("this is n")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
