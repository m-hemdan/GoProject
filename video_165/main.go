package main

import (
	"fmt"
)

func main() {
	cs := make(chan int) // if we create chan <- int that make send only
	// not reciever or priny

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
