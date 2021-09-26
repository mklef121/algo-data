package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value.")
		return
	}

	stringIndex := arguments[1]
	index, err := strconv.Atoi(stringIndex)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", index)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	if (len(aSlice) - 1) < index {
		fmt.Println("Cannot delete element", index)
		return
	}

	//Implementing first way of deleting an element as specified in array-and-slices.md
	aSlice = append(aSlice[:index], aSlice[index+1:]...)
	fmt.Println("After 1st deletion:", aSlice)

	// Implementing the second technique
	theSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// Replace element at index i with last element
	theSlice[index] = theSlice[len(theSlice)-1]
	// Remove last element
	theSlice = theSlice[:len(theSlice)-1]
	fmt.Println("After 2nd deletion:", theSlice)
}
