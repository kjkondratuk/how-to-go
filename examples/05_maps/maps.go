package main

import "fmt"

func main() {
	// create an empty map
	var myMap map[string]int

	if myMap == nil {
		fmt.Println("Map is empty!")
	}

	// need to allocate memory for the map or it'll panic
	myMap = make(map[string]int)

	// put some things in the map
	myMap["key1"] = 378
	myMap["key29"] = 27
	myMap["key23"] = 273

	fmt.Printf("map: %v\n", myMap)
}
