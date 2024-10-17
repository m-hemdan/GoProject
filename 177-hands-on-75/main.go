package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {

	fmt.Printf("The value of number %v and the type %T \n", a, a)
	fmt.Printf("The value of number %v and the type %T \n", b, b)
	fmt.Printf("The value of number %v and the type %T \n", c, c)
	fmt.Printf("The value of number %v and the type %T \n", d, d)

}
