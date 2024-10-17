package main

import "fmt"

func main() {

	fmt.Println(PrintSquare(square, 3))

}
func square(n int) int {
	return n * n
}
func PrintSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v", a, x)
}
