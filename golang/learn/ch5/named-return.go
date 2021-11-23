package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	fmt.Println(minMax(2, 3))

	// Functions that accept other functions as parameters
	data := []Grades{{"J.", "Lewis", 10}, {"M.", "Tsoukalos", 7},
		{"D.", "Tsoukalos", 8}, {"J.", "Lewis", 9}}
	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("It is sorted!")
	} else {
		fmt.Println("It is NOT sorted!")
	}
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}

	min = x
	max = y

	//Even without stating what to return the function returns the named variables called
	// min and max
	return

}

// Functions can return other functions

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}
	return func(k int) int {
		return k * k
	}
}
