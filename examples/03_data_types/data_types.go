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
	myInt := 234
	fmt.Printf("myInt is type: %T - value: %v\n", myInt, myInt)
	var myInt2 int64
	myInt2 = 9438759837459837
	fmt.Printf("myint2 is type: %T - value: %v\n", myInt2, myInt2)
	var myInt3 int32
	myInt3 = -943875983
	fmt.Printf("myint3 is type: %T - value: %v\n", myInt3, myInt3)
	var myInt4 uint
	myInt4 = 9438759837459837
	fmt.Printf("myint4 is type: %T - value: %v\n", myInt4, myInt4)
	var myInt5 uint64
	myInt5 = 943875983745983749
	fmt.Printf("myint5 is type: %T - value: %v\n", myInt5, myInt5)
	var myInt6 uint32
	myInt6 = 943875983
	fmt.Printf("myint6 is type: %T - value: %v\n", myInt6, myInt6)

	// float values
	myFloat := 1.237
	fmt.Printf("myFloat is type: %T - value: %v\n", myFloat, myFloat)
	myFloat2 := -1.29347293472932342
	fmt.Printf("myFloat2 is type: %T - value: %v\n", myFloat2, myFloat2)
	myFloat3 := 23974298374293742897342.23
	fmt.Printf("myFloat3 is type: %T - value: %v\n", myFloat3, myFloat3)

	// string values
	myString := "some test string"
	fmt.Printf("myString is type: %T - value: %v\n", myString, myString)
	myString2 := "🚀✅🥸🤩"
	fmt.Printf("myString2 is type: %T - value: %v\n", myString2, myString2)

	// first class functions
	myFunc := func() {
		fmt.Println("I'm a function stored in a variable.")
	}
	myFunc()
}
