package main

import "fmt"

type person struct {
	first  string
	last   string
	flavor []string
}

func main() {
	p1 := person{

		first:  "Marwa",
		last:   "hemdan",
		flavor: []string{"chocholate", "vanella", "banana"},
	}

	p2 := person{

		first:  "Marwa",
		last:   "hemdan",
		flavor: []string{"strawberry", "coconot", "mango"},
	}

	fmt.Printf(" the type is %T and the value is %v \n", p1, p1)
	fmt.Printf(" the type is %T and the value is %v \n", p2, p2)
	for _, v := range p1.flavor {
		fmt.Println(p1.first, "the value is ", v)
	}
	for _, v := range p2.flavor {
		fmt.Println(v)
	}

}
