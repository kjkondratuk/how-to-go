package main

import (
	"fmt"
	"strconv"
)

func main() {
	// call a function
	total, err := addNumbers("4", 3)
	if err != nil {
		fmt.Println("There was a problem adding 4 and 3")
	}
	fmt.Printf("First total was %d\n", total)

	total2, err2 := addNumbers("sj", 4)
	if err2 != nil {
		fmt.Println("There was a problem adding sj and 4")
	}
	fmt.Printf("Second total was %d\n", total2)

	// call an attached function
	s := someStruct{43}
	s.AddNumbers(3)
	s.AddNumbers(14, 34)

	fmt.Printf("Accumulator value: %d\n", )
}

func addNumbers(param1 string, param2 int) (int, error) {
	val, err := strconv.Atoi(param1)
	if err != nil {
		return 0, err
	}
	return val + param2, nil
}

type someStruct struct {
	value int
}

// attach a function to the struct

// AddNumbers : adds all the numbers specified to the struct value
func (s someStruct) AddNumbers(p2 ...int) int {
	for _, v := range p2 {
		s.value += v
	}
	return s.value
}
