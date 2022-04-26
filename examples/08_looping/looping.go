package main

import (
	"fmt"
	"time"
)

func main() {
	// loop forever (well, for a few seconds)
	for {
		fmt.Println("Looping...")
		time.Sleep(5 * time.Second)
		fmt.Println("Stopping loop...")
		break
	}

	// for each - ignore index
	str := []string{
		"a", "string", "array",
	}
	for _, s := range str {
		fmt.Printf("for each: %s\n", s)
	}

	// for each, multi-value
	mm := map[string]int{
		"a":      28,
		"string": 84,
		"value":  62,
	}
	for k, v := range mm {
		fmt.Printf("for each multi: key - %s : value - %d\n", k, v)
	}

	// fully qualified for-loop
	for i := 0; i < 8; i++ {
		fmt.Printf("fq for loop: index - %d\n", i)
	}
}
