package main

import "fmt"

func main() {
	trueValue := true
	anotherTrueValue := true
	aFalseValue := false

	// if statement
	if trueValue {
		fmt.Println("This is true")
		// else if
	} else if aFalseValue == anotherTrueValue || aFalseValue == trueValue {
		fmt.Println("This is false")
		// else
	} else {
		fmt.Println("Neither of the above are true")
	}

	// switch
	value := 72
	switch value {
	case 32:
		fmt.Printf("%d is equal to 32\n", value)
	case 72:
		fmt.Printf("%d is is equal to 72\n", value)
	default:
		fmt.Printf("%d is not equal to the handled numbers\n", value)
	}

	// type switch
	var myVal interface{}
	myVal = "test string"
	switch myVal.(type) {
	case string:
		fmt.Println("The value was a string")
	case int:
		fmt.Println("The value was an int")
	default:
		fmt.Println("The value was not of a known type")
	}

}
