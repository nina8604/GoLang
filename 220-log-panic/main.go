package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("name.txt")
	if err != nil {
		log.Panicln(err)
	}

}
