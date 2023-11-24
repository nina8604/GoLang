package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)

	// the blank idetifier to not use a returned value
	for _, value := range xs {
		fmt.Printf("%v\n", value)
	}

	fmt.Println("-----------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	fmt.Println("-----------")

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
}
