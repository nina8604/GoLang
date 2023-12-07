package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Name string
	Age  int
}

func main() {
	d1 := dog{"Hanse", 6}
	d2 := dog{"Sema", 4}

	pets := []dog{d1, d2}
	fmt.Println(pets)

	ps, err := json.Marshal(pets)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ps))
}
