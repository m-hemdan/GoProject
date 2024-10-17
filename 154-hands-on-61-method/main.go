package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("the name is ", p.first, "the age is ", p.age)
}
func main() {
	p := person{
		first: "Hisham",
		age:   37,
	}
	p.speak()
}
