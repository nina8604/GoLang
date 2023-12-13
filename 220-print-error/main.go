package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("name.txt")
	if err != nil {
		fmt.Println("error happened", err)
	}

}
