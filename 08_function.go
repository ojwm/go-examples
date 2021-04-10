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

func do(f func(int) int) int {
	return f(8)
}

func main() {
	names := []string{"Oli", "John", "Terry"}
	for _, name := range names {
		fmt.Print(hello(name))
	}
	fmt.Println("---------------------------")
	power2, power3 := power(4)
	fmt.Println(power2, power3)
	fmt.Println("---------------------------")
	invert := func(value int) int {
		return value * -1
	}
	increment := func(value int) int {
		return value + 1
	}
	fmt.Println(do(invert))
	fmt.Println(do(increment))
}
