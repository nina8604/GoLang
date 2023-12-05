package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)  // we see a value of x
	fmt.Println(&x) // we see an adress of x (pointer of x)

	i := 32
	p := &i         // 'p' is a pointer to 'i'
	fmt.Println(*p) // dereferencing 'p' gives you integer that 'p' points to
	*p = 21         // you can change a value that 'p' points to
	fmt.Println(i)  // the value of 'i' is now 21

}
