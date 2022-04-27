package main

import (
	"fmt"
	"github.com/kjkondratuk/how-to-go/examples/13_testing/stack"
)

type AStruct struct {
	SomeValue int
}

func main() {
	a := "some string"
	b := 273
	c := AStruct{SomeValue: 28}

	s := stack.NewStack()
	s.Push(a)
	s.Push(b)
	s.Push(c)

	fmt.Printf("Stack: %+v\n", s)

	// First value popped should be the last one on
	v1 := s.Pop().(AStruct)
	fmt.Printf("Popped: %+v\n", v1)
	fmt.Printf("Stack: %+v\n", s)

	// Second value popped should be the second one on, which was an integer
	v2 := s.Pop().(int)
	fmt.Printf("Popped: %d\n", v2)
	fmt.Printf("Stack: %+v\n", s)
}
