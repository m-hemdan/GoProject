package main

import "fmt"

func addA(a, b int) int {
	return a + b
}
func addB(a, b float64) float64 {
	return a + b
}

type myNum interface {
	~int | ~float64
}

func addT[T myNum](a, b T) T {
	return a + b
}

type newVar int

func main() {
	var n newVar = 42
	fmt.Println(addA(n, 2))
	fmt.Println(addB(1.2, 2.2))
	fmt.Println(addT(1.2, 2.3))

}
