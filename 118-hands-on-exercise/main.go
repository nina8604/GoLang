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
	fmt.Printf("%#v\n", x)

	for i, v := range x {
		fmt.Println(i)
		for j, vv := range v {
			fmt.Println(j, vv)
		}
	}

}
