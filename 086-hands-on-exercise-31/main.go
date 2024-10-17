package main

import "fmt"

func main() {
	m := map[string]int{
		"marwa":  33,
		"hisham": 37,
	}
	for k, v := range m {
		fmt.Printf("the name is %v and the age is %v", k, v)
	}
	fmt.Println("--------------------")
	age1 := m["marwa"]
	fmt.Println("The age of Bond", age1)
	if v, ok := m["marwa"]; ok {
		fmt.Println("There is a BOND lookup entry, and Bond's age is", v, ok)
	}

	age2 := m["Q"]
	fmt.Println("That age of Q", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int", v, ok)
	}
}
