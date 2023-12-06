package main

import "fmt"

func main() {
	i := 37
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(*&i)
}
