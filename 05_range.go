package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{43, 4, 567, 2, 645, 85, 4, 92, 3}
	// Find duplicate elements
	for i, element := range slice {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] == element {
				fmt.Println(element)
			}
		}
	}
}
