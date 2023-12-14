package main

import (
	"fmt"

	"github.com/nina8604/GoLang/tree/master/226-hands-on-exercise-documentation/dog"
)

func main() {
	ha := 28
	da := dog.Years(ha)
	fmt.Printf("human - %v years turn to dog years - %v", ha, da)
}
