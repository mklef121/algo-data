package main

import "fmt"

var sortArray = []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}

func main() {
	fmt.Println("Before the sorting", sortArray)
	insertionSort(sortArray)

	fmt.Println("\n\n after the sorting", sortArray)
}

func insertionSort(data []int) {
	count := len(data)
	for i := 0; i < count; i++ {
		current := data[i]
		j := i - 1

		/*

			while(j >= 0 && current < data[j]){

			}
		*/
		for j >= 0 && current < data[j] {
			data[j+1] = data[j]

			j--
		}

		data[j+1] = current
	}
}
