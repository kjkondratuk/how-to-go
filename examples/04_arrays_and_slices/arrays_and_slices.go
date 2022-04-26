package main

import "fmt"

func main() {
	// create an array of integers
	integers := []int{
		12,
		83,
		23,
		18,
		87,
	}
	idx := 3
	// select the number at a given index
	fmt.Printf("index: %d - value: %d\n", idx, integers[idx])

	// strings are arrays too (of runes)
	myStr := "some test string"
	idx = 5
	fmt.Printf("index: %d - value: %s\n", idx, string(myStr[idx]))

	// get a slice of multiple values
	idxE := 9
	fmt.Printf("index: %d to %d - value: %s\n", idx, idxE, myStr[idx:idxE])

	// strings are immutable (uncomment to see)
	//myStr[idx] = 'r'

	// Create an array of a certain size full of default values
	emptyArr := make([]int, 10)
	fmt.Printf("empty array: %v\n", emptyArr)

	// set one element
	emptyArr[7] = 27
	fmt.Printf("empty array w/update: %v\n", emptyArr)
}
