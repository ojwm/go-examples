package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3}
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	fmt.Println("Array sum:", sum)
	array2d := [][]int{{1, 2}, {3, 4}, {5, 6}}
	sum = 0
	for i := 0; i < len(array2d); i++ {
		for j := 0; j < len(array2d[i]); j++ {
			sum += array2d[i][j]
		}
	}
	fmt.Println("Array2D sum:", sum)
}
