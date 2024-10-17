package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	width  float64
}
type CIRCLE struct {
	radius float64
}

func (S SQUARE) area() float64 {
	return S.length * S.width
}

func (C CIRCLE) area() float64 {
	return math.Pi * math.Pow(C.radius, 2)

}

type Shape interface {
	area() float64
}

func info(Sh Shape) {
	fmt.Println("the area is", Sh.area())
}
func main() {
	S := SQUARE{
		length: 5,
		width:  8,
	}
	C := CIRCLE{
		radius: 4,
	}

	info(S)
	info(C)
}
