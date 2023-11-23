package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("the key is %v\t the value is %v\n", k, v)
	}

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
}
