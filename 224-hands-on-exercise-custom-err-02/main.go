package main

import (
	"fmt"
)

type customErr struct {
	message string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("this is err: %v", ce.message)
}

func main() {
	ce1 := customErr{
		message: "the custom err",
	}

	foo(ce1)

}

func foo(err error) {
	// fmt.Println("foo ran - ", err)
	// conversial and
	fmt.Println("foo ran - ", err, "\n", err.(customErr).message)
}
