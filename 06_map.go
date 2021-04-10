package main

import (
	"fmt"
)

func main() {
	var fruits map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println("pear:", fruits["pear"])
	fruits["pear"] = 2
	fmt.Println("pear:", fruits["pear"])
	delete(fruits, "apple")
	fmt.Println(fruits)
	vegetables := make(map[string]int)
	fmt.Println(vegetables)
	vegetables["carrot"] = 8
	fmt.Println(vegetables)
	value, present := vegetables["carrot"]
	if present {
		fmt.Println("carrot:", value)
	}
	value2, present2 := vegetables["cucumber"]
	if present2 {
		fmt.Println("cucumber:", value2)
	} else {
		fmt.Println("cucumber: none")
	}
}
