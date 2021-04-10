package main

import (
	"fmt"
)

func hello(name string) string {
	return fmt.Sprintf("Hello %s!\n", name)
}

func power(value int) (int, int) {
	// defer something to after the return (useful for cleanup)
	defer fmt.Println("power function end")
	fmt.Println("power function start")
	return value * value, value * value * value
}

func main() {
	names := []string{"Oli", "John", "Terry"}
	for _, name := range names {
		fmt.Print(hello(name))
	}
	power2, power3 := power(4)
	fmt.Println(power2, power3)
}
