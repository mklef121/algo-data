package main

import (
	"fmt"
)

var sortArray = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

func main() {
	fmt.Println("Before the sorting", sortArray)
	bubbleSort(sortArray)

	fmt.Println("\n\n after the sorting", sortArray)

}

func bubbleSort(data []int) {
	count := len(data)
	for i := 0; i < count; i++ {

		for j := 0; j < count; j++ {
			if data[j] > data[i] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

}
