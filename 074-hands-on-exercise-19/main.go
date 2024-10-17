package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("it is the main function \n")
}
func init() {
	fmt.Print("it is the init function\n")

	x := rand.Intn(400)
	fmt.Printf("the random number is %v and the type is %T \n", x, x)
	if x >= 0 && x <= 100 {
		fmt.Println("print between 0 and 100")
	}
	if x >= 101 && x <= 200 {
		fmt.Println("print between 101 and 200")
	}
	if x >= 201 && x <= 250 {
		fmt.Println("print between 201 and 250")
	}
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

	// exercise 75

}
