package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //4
	fmt.Println(f()) //5
	fmt.Println(f()) //6

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x

	}
}
