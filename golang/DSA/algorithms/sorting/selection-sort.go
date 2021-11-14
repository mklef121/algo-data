package main

import "fmt"

var sortArray = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

func main() {
	fmt.Println("Before the sorting", sortArray)
	selectionSort(sortArray)

	fmt.Println("\n\n after the sorting", sortArray)
}

func selectionSort(data []int) {
	count := len(data)
	for i := 0; i < count; i++ {
		minIndex := i

		for j := i + 1; j < count; j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}

		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
