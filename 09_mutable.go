package main

import (
	"fmt"
)

func main() {
	// Literal data types are immutable
	var x int = 5
	y := x
	y = 9
	// x is not bound to y
	fmt.Println(x, y)
	fmt.Println("---------------------------")
	// Slice is a mutable data type
	// Slice is a pointer to an array in RAM
	var a []int = []int{1, 2, 3, 4, 5}
	b := a
	b[0] = 999
	// a is bound to b
	fmt.Println(a, b)
	fmt.Println("---------------------------")
	// Map is a mutable data type
	// Map is a pointer to an array in RAM
	var p map[string]int = map[string]int{"car": 6, "truck": 3}
	q := p
	delete(q, "car")
	// p is bound to q
	fmt.Println(p, q)
	fmt.Println("---------------------------")
	// Array is a mutable data type
	// Array is not a pointer
	var u [3]int = [3]int{1, 2, 3}
	v := u
	v[0] = 999
	fmt.Println(u, v)
	fmt.Println("---------------------------")
	// Use a function to demonstrate modifying a slice
	// without using a return value
	var slice []int = []int{1, 2, 3}
	slice2 := slice
	fmt.Println("Before: ", slice)
	fmt.Println("Before2:", slice2)
	add := func(s []int) {
		s[0] += 1000
	}
	add(slice2)
	fmt.Println("After: ", slice)
	fmt.Println("After2:", slice2)
}
