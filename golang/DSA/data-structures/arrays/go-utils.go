package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	arguments := os.Args
	strToReverse := "hello pumpkin"
	if len(arguments) != 1 {
		strToReverse = arguments[1]
	}

	reversedString := reverseString(strToReverse)

	fmt.Println("reversedString: ", reversedString)

	fmt.Println("Merging Sorted arrays: ", mergeSortedArrays([]int{2, 3, 4, 5, 20, 76}, []int{1, 8, 10, 22, 45}))

}

func reverseString(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	var newString []byte

	for i := length - 1; i >= 0; i-- {
		// fmt.Println(i)
		newString = append(newString, s[i])
	}

	return string(newString)
}

func mergeSortedArrays(arr1 []int, arr2 []int) []int {

	if len(arr2) == 0 {
		return arr1
	}

	if len(arr1) == 0 {
		return arr2
	}

	finalArray := make([]int, len(arr1)+len(arr2))

	arr1_iter := 0
	arr2_iter := 0

	arr1Current := arr1[arr1_iter]
	arr2Current := arr2[arr2_iter]
	rotate := 0

	l1 := len(arr1)
	l2 := len(arr2)
	totalRotate := 0
	for arr1_iter < l1 || arr2_iter < l2 {

		fmt.Println("rotate", rotate, "arr1_iter", arr1_iter, "arr2_iter", arr2_iter, finalArray)
		rotate++
		if arr1Current < arr2Current {
			arr1_iter++
			finalArray[totalRotate] = arr1Current
			if arr1_iter < l1 {
				arr1Current = arr1[arr1_iter]
			}
		} else {
			arr2_iter++
			finalArray[totalRotate] = arr2Current
			if arr2_iter < l2 {
				arr2Current = arr2[arr2_iter]
			} else {
				arr2Current = math.MaxInt32
			}
		}

		totalRotate++
	}

	// for arr1_iter < len(arr1) || arr2_iter < len(arr2) {

	// 	if arr1Current > arr2Current {
	// 		finalArray = append(finalArray, arr1Current)
	// 		if arr1_iter >= len(arr1) {
	// 			continue
	// 		}
	// 		arr1Current = arr1[arr1_iter]
	// 		arr1_iter++
	// 	} else {
	// 		finalArray = append(finalArray, arr2Current)
	// 		arr2Current = arr2[arr2_iter]
	// 		arr2_iter++
	// 	}

	// 	continue

	// }

	return finalArray

}
