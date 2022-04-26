package main

import "fmt"

type MyType interface {
	IsLess(i int) bool
}

type MyOtherType interface {
	MyType
	IsMore(i int) bool
}

// A implements only MyType
type A struct {
	value int
}

func (a *A) IsLess(i int) bool {
	return a.value < i
}

// B implements both MyType and MyOtherType
type B struct {
	value int
}

func (b *B) IsLess(i int) bool {
	return b.value < i
}

func (b *B) IsMore(i int) bool {
	return b.value > i
}

func main() {
	a := A{value: 21}
	b := B{value: 47}
	vals := []int{12, 67}

	// make a generic array of objects that implement IsLess
	arr := []MyType{
		&a, &b,
	}

	// can do type switching on interfaces
	for _, v := range arr {
		switch v.(type) {
		case *A:
			fmt.Printf("%v is an A\n", v)
		case *B:
			fmt.Printf("%v is a B\n", v)
		default:
			fmt.Printf("%v is unknown\n", v)
		}
	}

	// methods available are limited to the scoped type (MyType for arr)
	for i, v := range arr {
		if v.IsLess(vals[i]) {
			fmt.Printf("%d is Less than %d\n", v, vals[i])
		} else {
			fmt.Printf("%d is NOT less than %d\n", v, vals[i])
		}
	}

	for i, v := range arr {
		if i == 0 {
			// use a type assertion to get access to A's value
			fmt.Printf("Value of A is: %d\n", v.(*A).value)
		}
		if i == 1 {
			// use type assertion to compare the values of A and B
			fmt.Printf("%v is more than %v: %v", v, arr[0], v.(*B).IsMore(arr[0].(*A).value))
		}
	}

}
