package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Janny", "Moneypenny", "Guiness", "Wolvarine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)

}
