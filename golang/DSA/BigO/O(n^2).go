package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	a := []int{1, 3, 4, 5, 6, 7, 10, 60, 23, 17}
	b := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(b)

	// array900 := make([]string, 90000)

	formPairs(a)
	// loopArray(array900, "Large")

}

//This function runs with a linear BigO notation **O(n * n) =>> O(n^2)** Quadratic time
func formPairs(array []int) {
	count := len(array)
	// fmt.Println("Count is ", count)
	beginTIme := time.Now().UnixNano()
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			fmt.Println(strconv.Itoa(array[i]) + " " + strconv.Itoa(array[j]))
		}
	}

	endTime := time.Now().UnixNano()

	// time.Now().Unix()
	timeDiff := float64(endTime - beginTIme)
	var endMilliSeconds float64 = timeDiff / 1000000
	fmt.Println(beginTIme, endTime)
	fmt.Println("The completion time for the array is : ", endMilliSeconds, " Milli Seconds \n -->in Nano seconds is ", timeDiff)
}
