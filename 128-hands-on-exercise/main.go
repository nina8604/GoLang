package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors string
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Mercedess",
		model: "v100",
		doors: "four",
		color: "red",
	}
	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Hundai",
		model: "x101",
		doors: "two",
		color: "black",
	}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.engine, v1.make, v1.model, v1.doors, v1.color)
	fmt.Println(v2.engine, v2.make, v2.model, v2.doors, v2.color)

}
