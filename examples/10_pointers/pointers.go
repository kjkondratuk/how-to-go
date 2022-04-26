package main

import "fmt"

func main() {
	x := "my string"
	// create a pointer to x
	y := &x
	// create a copy of x
	n := x

	fmt.Printf("x: %s - y: %v (%s) - n: %s\n", x, y, *y, n)
	// update n - no change in x
	n = "wf0hwef"
	fmt.Printf("x is still: %s\n", x)

	// update value at y - x should change too
	*y = "my other string"
	fmt.Printf("x is now: %s\n", x)

	// an indexed reference to an array element is a pointer
	myInts := []int{12, 46, 87, 21, 16}
	myInts[1] = 99

	fmt.Printf("myInts: %v\n", myInts)

	// Create a reference to myInts
	myOtherInts := myInts
	myOtherInts[3] = 99

	fmt.Printf("myInts: %v\n", myInts)

	// slicing maintains reference to the original array
	mySlicedInts := myInts[2:5]
	mySlicedInts[2] = 99

	fmt.Printf("myInts updated slice: %v\n", myInts)

	// strings are still immutable
	//myString := "abc123"
	//mySubstring := myString[2:4]
	//mySubstring[0] = 'X'
}
