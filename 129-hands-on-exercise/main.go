package main

import "fmt"

func main() {

	an := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Nina",
		friends: map[string]int{
			"Inna":   39,
			"Nastia": 38,
		},
		favDrinks: []string{"Martini", "Beer"},z
	}

	fmt.Println(an)
	fmt.Printf("%#v", an)

}
