package main

import (
	"fmt"
)

func main() {
	// Use empty struct to implement sets
	fruits := make(map[string]struct{})
	fruits["apple"] = struct{}{}
	fruits["banana"] = struct{}{}
	fmt.Println(fruits)
	_, present := fruits["banana"]
	if present {
		fmt.Println("banana")
	}
	// Optionally make empty struct a variable
	// to avoid using struct{}{}, which may be confusing
	var empty struct{}
	vegetables := make(map[string]struct{})
	vegetables["carrot"] = empty
	vegetables["potato"] = empty
	fmt.Println(vegetables)
	_, present2 := vegetables["potato"]
	if present2 {
		fmt.Println("potato")
	}
}
