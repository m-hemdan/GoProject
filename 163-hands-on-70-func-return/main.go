package main

import "fmt"

func main() {
	f := x()
	fmt.Println(f())
}

func x() func() int {
	return func() int {
		return 42
	}
}
