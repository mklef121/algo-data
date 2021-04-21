package main

import (
	"fmt"
	"time"
)

func main() {
	a := []string{"nemo", "Ultimate", "coding", "interview", "bootcamp", "Get",
		"more", "job", "offers", "negotiate", "a", "raise", "Everything",
		"you", "need", "to", "get", "the", "job", "you", "want"}

	// array900 := make([]string, 90000)

	loopArray(a, "Small")
	// loopArray(array900, "Large")

}

//This function runs with a linear BigO notation **O(n)**
func loopArray(array []string, description string) {
	count := len(array)
	// fmt.Println("Count is ", count)
	beginTIme := time.Now().UnixNano()
	for i := 0; i < count; i++ {
		if array[i] == "nemo" {
			// one notable factor was how long the `printLn` command was delaying the processing time

			// fmt.Println("Found Nemo at index :", i)
		}
	}

	endTime := time.Now().UnixNano()

	// time.Now().Unix()
	timeDiff := float64(endTime - beginTIme)
	var endMilliSeconds float64 = timeDiff / 1000000
	fmt.Println(beginTIme, endTime)
	fmt.Println("The completion time for "+description+" array is : ", endMilliSeconds, " Milli Seconds \n -->in Nano seconds is ", timeDiff)
}
