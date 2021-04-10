package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array", array)
	// Slice starts with the value at index
	// Slice ends with the value at index-1
	var slice []int = array[1:3]
	fmt.Println("Slice", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))
	fmt.Printf("Array type: %T\n", array)
	fmt.Printf("Slice type: %T\n", slice)
	fmt.Println("-------------------")
	var newSlice []int = []int{5, 6, 7, 8, 9}
	extendedSlice := append(newSlice, 10)
	fmt.Println("New slice:", newSlice)
	fmt.Println("Extended slice:", extendedSlice)
	makeSlice := make([]int, 5)
	fmt.Println("Make slice:", makeSlice)
}
