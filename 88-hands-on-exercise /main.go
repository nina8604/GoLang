package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("the key is %v\t the value is %v\n", k, v)
	}
	fmt.Println("--------")

	age1 := m["James"]
	fmt.Println("The age of Bond", age1)

	if v, ok := m["James"]; ok {
		fmt.Println("This is a Bond lookup entery and Bond's age is", v)
	}

	age2 := m["Q"]
	fmt.Println("the age of Q", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("This is no Q and here is zero value of an int", v)
	}
	fmt.Println("--------")
	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("the iteration is %v\t total count %v\t x is %v\n", i, c, x)
			c++
		}
	}

}
