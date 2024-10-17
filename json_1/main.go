package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type colorStruct struct {
		Id     int
		Name   string
		Colors []string
	}
	group := colorStruct{
		Id:     1,
		Name:   "red",
		Colors: []string{"a", "b", "c"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)

	}
	os.Stdout.Write(b)
}
