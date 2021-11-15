package main

import "fmt"

func main() {

	// left := []int{1, 2, 5, 6, 8}
	// right := []int{0, 3, 4, 7, 9, 20}
	toSortArray := []int{89, 55, 0, 3, 4, 1, 2, 5, 1000, 6, 8, 7, 9, 20}
	// fmt.Println(mergeTwoOrdered(left, right))

	fmt.Println(mergeSort(toSortArray))
}

func mergeSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2

	left := arr[:mid]
	right := arr[mid:]

	return mergeTwoOrdered(mergeSort(left), mergeSort(right))
}

func mergeTwoOrdered(left []int, right []int) []int {

	// fmt.Println(len(left), len(right), "The lengths")
	if len(right) == 0 {
		return left
	}

	if len(left) == 0 {
		return right
	}

	merger := make([]int, len(left)+len(right))
	mergerFillIndex := 0

	//The number of times we have iterated over left
	leftItered := 0
	rightItered := 0

	leftVal := left[leftItered]
	rightVal := right[rightItered]

	for leftItered < len(left) || rightItered < len(right) {

		if leftVal <= rightVal && leftItered < len(left) || rightItered >= len(right) {
			merger[mergerFillIndex] = leftVal
			mergerFillIndex++
			leftItered++

			if leftItered < len(left) {
				leftVal = left[leftItered]
			}

		} else {
			// fmt.Println("The mergerFillIndex", mergerFillIndex, rightItered)
			merger[mergerFillIndex] = rightVal
			mergerFillIndex++
			rightItered++
			if rightItered < len(right) {
				rightVal = right[rightItered]
			}
		}
	}

	return merger
}
