package main

import "fmt"

func main() {
	str := []string{"Come", "Go", "Hello"}

	//This is an O(1), No matter how many elements arte in the array
	// Just a single operation is performed
	res := pickArrayValue(str, 2)
	fmt.Println(res)
}

func pickArrayValue(arr []string, index int) string {
	return arr[index]
}
