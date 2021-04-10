package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 110; i++ {
		if i > 0 && i%9 == 0 {
			var size string
			stringValue := strconv.Itoa(i)
			switch len(stringValue) {
			case 1:
				size = "S"
			case 2:
				size = "M"
			case 3:
				size = "L"
			}
			fmt.Printf("%s %3d\n", size, i)
		}
	}
}
