package main

import "fmt"

func main() {
	// x := map[string][]string{
	// 	"bond_james":       []string{"shaken, not stirred", "martinis", "fast cars"},
	// 	"moneypenny_jenny": []string{"intelligence", "literature", "computer science"},
	// 	"no_dr":            []string{"cats", "ice cream", "sunsets"},
	// }

	x := make(map[string][]string)
	x["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	x["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	x["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	x["flaming_ian"] = []string{"steaks", "cigars", "espionage"}

	for i, v := range x {
		fmt.Println(i)
		for j, vv := range v {
			fmt.Println(j, vv)
		}
	}
	fmt.Println("------delete record-------")

	delete(x, "no_dr")

	fmt.Println("------delete record-------")

	for i, v := range x {
		fmt.Println(i)
		for j, vv := range v {
			fmt.Println(j, vv)
		}
	}

}
