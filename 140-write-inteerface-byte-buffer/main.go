package main

import (
	"bytes"
	"fmt"
)

// type person struct{
// 	first string
// }

// func (p person) WriteOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	// //create a file
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer f.Close()

	// s := []byte("Hello gopher!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("gopher!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("It is Thursday\n")
	fmt.Println(b.String())

	b.Write([]byte("Happy happy"))
	fmt.Println(b.String())

}
