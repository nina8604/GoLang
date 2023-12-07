package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 47, 21, 15, 98, 73, 14, 2, 8, 1}
	xs := []string{"James", "Q", "M", "MOneypenny", "Dr.No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
