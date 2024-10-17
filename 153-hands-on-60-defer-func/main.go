package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println(bar())
	fmt.Println(foo())

}
func foo() int {
	return 5
}
func bar() string {
	return "Marwa"
}
