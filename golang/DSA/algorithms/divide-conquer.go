package main

import (
	"fmt"
)

/*
Bitcoin withing a period went up and down in prices on a chart (y- axis) in the following manner

-2,1,-3,4,-1,2,1,-5,4

# In this period of nine days, what can be the maximum amount of profit you can make

# This is the maximum sub array proble

we can use

- Divide and conquer
- Greedy algorithms
- Dynamic Programming

#Finding a solution
put prices in an array [-2,1,-3,4,-1,2,1,-5,4]

The max subarray will either be in the leftmost part or rightmost part when the array is divided.

or Is in the middle crossing into the left and right of the array.

We will continously divide each wing until we get to the last
100000

1000010001

var count int = 0
var max int = 0
if i = 1

	if count > max
		max = count

	count = 0

	startCount = true

if i = 0 && startCount

	count++
*/
func main() {
	// hi := 5 / 2
	var (
	// 8-bit bytes : "byte array".
	// [0,1,2 ...,255]
	// me string
	// byte is an alias for uint8 and is equivalent to uint8 in all ways
	// it's unsigned 8 bit integer value
	// Range: 0 through 255.
	// you []byte

	// jg rune = '😇'
	)

	reverseString([]byte{'h', 'e', 'l', 'l', 'o', 'a', '!'})

	array := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	arrLength := len(array)
	leftLen := len(array) / 2
	rightLen := arrLength - leftLen

	fmt.Println((array), leftLen, rightLen, array[:leftLen], array[leftLen:])

	sideMax(array, 0, arrLength-1)
}

func reverseString(s []byte) {
	fmt.Println(s, " :::these bytes rep =>", string(s))
}

func sideMax(arr []int, start int, end int) int {

	if start == end {
		return arr[start]
	} else {
		mid := (start + end) / 2
		leftMax := sideMax(arr, start, mid)
		rightMax := sideMax(arr, mid+1, end)

		fmt.Println("The end", leftMax, rightMax)
	}

	return 1
}

// max
