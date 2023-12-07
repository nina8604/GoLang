package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(bs)

	loginPassword1 := "password1234"
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1))
	if err != nil {
		fmt.Println("you can't login")
		return
	}
	fmt.Println("we are logged in")

}
