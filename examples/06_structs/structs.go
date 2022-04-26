package main

import "fmt"

type customer struct {
	customerId string
	score      int
	Name       string
}

func main() {
	// An anoymous struct with positional attributes
	myStruct := struct {
		customerId string
		score      int
		name       string
	}{
		"92fh92",
		87,
		"Jerry Lee",
	}

	fmt.Printf("anonymous struct: %v\n", myStruct)

	// Access a struct attribute
	fmt.Printf("customerId: %s\n", myStruct.name)

	// Instantiate a named struct, with named attributes
	myCustomer := customer{
		customerId: "w337gdf3",
		score:      49,
		Name:       "Sharon Weisse",
	}

	fmt.Printf("named struct: %v\n", myCustomer)
}
