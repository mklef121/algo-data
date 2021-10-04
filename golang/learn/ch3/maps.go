package main

import "fmt"

func main() {
	record := map[string]int{
		"key1": 4,
		"key2": 123,
	}

	rec, ok := record["key1"]

	fmt.Println(rec, ok)
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)

	//this makes this map a nil map
	aMap = nil

	fmt.Println("a nil aMap:", aMap)

	// This will fail
	// aMap["test"] = 1

	// Iterating over maps

	theMap := make(map[string]string)

	theMap["123"] = "456"
	theMap["key"] = "A value"

	for key, value := range theMap {
		fmt.Println("key:", key, "value:", value)
	}

}
