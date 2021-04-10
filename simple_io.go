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
	fmt.Print("Enter your age: ")
	scanner.Scan()
	input = scanner.Text()
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("'%s' is not an integer.\n", input)
	} else {
		var template, message string = "You can drive %s.\n", ""
		if age >= 18 {
			message = "alone"
		} else if age >= 17 {
			message = "with a fully licensed passenger"
		} else {
			message = "when you're older"
		}
		fmt.Printf(template, message)
	}
}
