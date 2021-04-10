package main

import (
	"fmt"
)

func change(s string) {
	s = "modified"
}

func changeWithPointer(s *string) {
	*s = "modified with pointer"
}

func main() {
	var x int = 7
	y := &x
	fmt.Println("Value:", x, "Pointer:", y)
	// Modify the data at the pointer location
	*y = 8
	fmt.Println("Value:", x, "Pointer:", y)
	fmt.Println("---------------------------")
	var value string = "default"
	fmt.Println(value)
	// This won't change, as literals are immutable
	change(value)
	fmt.Println(value)
	// This will change, as the pointer to the value is being used
	changeWithPointer(&value)
	fmt.Println(value)
	fmt.Println("---------------------------")
	var thing string = "default"
	var p *string = &thing
	// p is the pointer to the value of thing
	fmt.Println(p)
	// *p is the value of thing
	fmt.Println(*p)
	// &p is the pointer to the value of p
	// which is the pointer to the value of thing
	fmt.Println(&p)
}
