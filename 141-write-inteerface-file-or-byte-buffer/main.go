package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) WriteOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {

	p := person{
		first: "James",
	}
	// //create a file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer
	p.WriteOut(f)
	p.WriteOut(&b)
	fmt.Println(b.String())

}
