package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*	for i := 1; i < 99; i++ {
		fmt.Println(i)
	}*/
	for i := 0; i < 42; i++ {

		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("the value of %v the number x is %v \n", i, x)
		case 1:
			fmt.Printf("the value of %v the number x is %v \n", i, x)
		case 2:
			fmt.Printf("the value of %v the number x is %v \n", i, x)
		case 3:
			fmt.Printf("the value of %v the number x is %v \n", i, x)
		case 4:
			fmt.Printf("the value of %v the number x is %v \n", i, x)

		}
	}
	j := 0
	for j < 10 {
		fmt.Printf("the value of j %v \n", j)
		j++
	}
}
