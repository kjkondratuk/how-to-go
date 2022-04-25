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
}
