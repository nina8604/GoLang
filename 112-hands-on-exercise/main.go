package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "I'm 008"}

	x := [][]string{jb, mm}
	for i, v := range x {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
