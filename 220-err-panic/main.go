package main

import (
	"os"
)

func main() {
	_, err := os.Open("name.txt")
	if err != nil {
		panic(err)
	}

}
