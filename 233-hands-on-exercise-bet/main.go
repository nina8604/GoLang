package main

import (
	"fmt"
	"mymodule/233-hands-on-exercise-bet/quote"
	"mymodule/233-hands-on-exercise-bet/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
