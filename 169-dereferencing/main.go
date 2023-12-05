package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)  // we see a value of x
	fmt.Println(&x) // we see an adress of x (pointer of x)

	y := &x
	fmt.Printf("%v\t %T\t", y, y)
	fmt.Println(*y)

}
