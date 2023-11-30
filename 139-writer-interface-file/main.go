package main

import (
	"log"
	"os"
)

// type person struct{
// 	first string
// }

// func (p person) WriteOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	//create a file
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello gopher!")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

}
