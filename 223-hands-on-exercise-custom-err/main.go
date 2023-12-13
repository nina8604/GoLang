package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		// fmt.Println(err)
		// return
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, errors.New("the err in the json")
		return []byte{}, errors.New(fmt.Sprintf("the err in the json: %v", err))
		// return []byte{}, fmt.Errorf("the err in the json: %v", err)
		// return []byte{}, fmt.Println("the err in the json", err) // it doesn't work
	}
	return bs, nil
}
