package main

import "fmt"

func main() {
	fmt.Println(Paradise("xxx"))
}
func Paradise(log string) string {
	return fmt.Sprintf("this idea is %s", log)

}
