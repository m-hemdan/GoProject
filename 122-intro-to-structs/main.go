package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Marwa",
		last:  "hemdan",
		age:   32,
	}

	fmt.Printf("defination %v", p1)

}
