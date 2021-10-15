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

	fillme := twoSum([]int{-3, 4, 3, 90}, 0)

	fmt.Println("twoSum: ", fillme)

	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

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

	l1 := len(arr1)
	l2 := len(arr2)
	totalRotate := 0
	for arr1_iter < l1 || arr2_iter < l2 {

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

	return finalArray

}

func twoSum(nums []int, target int) []int {
	var dest []int
	holdMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		current := nums[i]

		remainder := target - current
		val, exists := holdMap[current]
		fmt.Println(current, remainder, exists, holdMap, "The holds")
		if exists {
			dest = append(dest, val, i)
			break
		}

		holdMap[remainder] = i
		// fmt.Println(holdMap, "The holds")

	}
	return dest
}

func twoSumCopy(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{0, 0}
}

func maxSubArray(nums []int) int {

	largest := 0
	largestSumArrays := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if i == 0 {
			largest = sum
			largestSumArrays[0] = i
		}

		for j := i; j < len(nums); j++ {

			if j != i {
				sum = sum + nums[j]
			}

			if sum > largest {
				largest = sum
				largestSumArrays[0] = i
				largestSumArrays[1] = j
			}
		}
	}

	return largest
}
