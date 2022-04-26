package main

import "fmt"

// customer is a struct.  Structs are similar to the C concept of structs and simply denote
// the layout of a region of memory.
type customer struct {
	customerId string
	score      int
	Name       string
}

type customer2 struct {
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

	// cast a struct to another with the same layout (doesn't work if they are different)
	otherCustomer := customer2(myCustomer)

	fmt.Printf("other customer: %v\n", otherCustomer)
}
