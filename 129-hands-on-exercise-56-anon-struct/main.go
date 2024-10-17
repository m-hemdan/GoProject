package main

import "fmt"

// type food struct {
// 	first    string
// 	friends  map[string]int
// 	favDrink []string
// }

func main() {
	food1 := struct {
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first:    "name1",
		friends:  map[string]int{"key1": 1, "key2": 2},
		favDrink: []string{"soda", "pepsi", "juice"},
	}
	fmt.Println(food1.first)
	for k, v := range food1.friends {
		fmt.Println("the values of key", k, "is ", v)
	}
	for _, v := range food1.favDrink {
		fmt.Println("the values of key is ", v)
	}

}
