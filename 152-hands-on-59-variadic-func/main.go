package main

import "fmt"

func main() {
	xx := []int{1, 2, 3, 4, 5, 6, 7}
	sum := foo(xx...)
	fmt.Println("the sum is ", sum)
	sumY := bar(xx)
	fmt.Println("the sum is ", sumY)

}
func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v

	}
	return sum
}
func bar(y []int) int {
	sum := 0
	for _, v := range y {
		sum += v

	}
	return sum

}
