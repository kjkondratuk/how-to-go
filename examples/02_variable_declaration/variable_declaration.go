package main

import (
	"fmt"
	"strings"
)

func main() {
	// Simple assignment, implicit type
	message := "Hello, world!"
	fmt.Println(message)

	// Simple assignment, explicit type
	var message2 string
	message2 = "Hello again, world!"
	fmt.Println(message2)

	// Multiple assignment, implicit type
	hello, world := "Hello!  ", "To the world!"
	parts := []string{hello, world} // make an array of strings
	fmt.Println(strings.Join(parts, ""))

	// You can assign to an underscore when you don't care about the result of an expression
	// the go compiler will yell at you if you don't use a variable you've declared
	_ = "my string"
}
