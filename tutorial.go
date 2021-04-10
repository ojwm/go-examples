package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input string = scanner.Text()
	fmt.Printf("Hello %s!\n", input)
	fmt.Print("Enter an integer: ")
	scanner.Scan()
	input = scanner.Text()
	intValue, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("'%s' is not an integer.\n", input)
	} else {
		var result bool = intValue > 10
		fmt.Printf("%s is greater than 10? %t\n", input, result)
	}
}
