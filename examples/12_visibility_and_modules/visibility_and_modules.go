package main

import (
	"fmt"
	"github.com/kjkondratuk/how-to-go/examples/12_public_vs_private/test_module"
)

func main() {
	// notice you can't access the private value here to set it when the struct is intiialized
	s := test_module.TestStruct{PublicValue: "some public value"}

	// printing the struct shows the private value has nothing, so how do we set it?
	fmt.Printf("struct: %+v\n", s)

	// Let's try using an accessor method
	s.SetPrivateValue("some other value")
	fmt.Printf("struct with private value set: %+v\n", s)
}
