package main

import (
	"fmt"
)

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

	E1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "make1",
		model: "model1",
		doors: "door1",
		color: "color1",
	}
	E2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "make2",
		model: "model2",
		doors: "door2",
		color: "color2",
	}
	fmt.Println("the value of e1 is ", E1, " \n the value of E2 is ", E2)
	fmt.Println(E1.make, E1.model, E1.doors, E1.color, E1.electric)
	fmt.Println(E2.make, E2.model, E2.doors, E2.color, E2.electric)

}
