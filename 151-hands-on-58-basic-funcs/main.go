package main

import "fmt"

func main() {

	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {

	return 5
}
func bar() (string, int) {

	return "Hello", 5
}
