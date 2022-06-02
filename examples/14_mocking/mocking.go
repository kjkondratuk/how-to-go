package main

import (
	"fmt"
	"github.com/kjkondratuk/how-to-go/examples/14_mocking/service"
	"github.com/kjkondratuk/how-to-go/examples/14_mocking/slow_service"
)

func main() {
	// create a new service with a new slow service
	s := service.New(slow_service.New())

	// create a list of integers to apply a power to
	integers := []int{12, 53, 23, 7, 2, 85, 298, 83, 21}
	fmt.Printf("original values: %v\n", integers)
	newValues := s.ApplyPower(integers)
	fmt.Printf("new values: %v\n", newValues)
}
