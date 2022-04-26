package main

import "fmt"

func main() {
	// boolean values
	boolVal := true
	boolVal2 := false
	boolVal3 := true

	if boolVal == boolVal3 {
		fmt.Printf("%v equals %v\n", boolVal, boolVal3)
	}

	if boolVal2 == boolVal3 {
		fmt.Printf("%v should NOT equal %v\n", boolVal2, boolVal3)
	}

	// integer values
	myInt := 1
	fmt.Printf("myInt is type: %T\n", myInt)

	var myInt2 int64
	myInt2 = 34
	fmt.Printf("myint2 is type: %T\n", myInt2)

}
