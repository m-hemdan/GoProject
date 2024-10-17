package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(400)
	switch {
	case x <= 100:
		fmt.Println("print between 0 and 100")

	case x >= 101 && x <= 200:
		{
			fmt.Println("print between 101 and 200")
		}
	case x >= 201 && x <= 250:
		{
			fmt.Println("print between 201 and 250")
		}
	default:
		fmt.Println("this was more than 2")

	}
}
