package main

import "fmt"

type youngin interface {
	walk()
	run()
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}
func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type dog struct {
	first string
}

func youngRun(y youngin) {
	y.run()
}

func main() {

	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)

}
func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

// uncomment this code to run it

/*



 */
