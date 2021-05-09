package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	//Merge two sorted arrays

	s1 := []int{1, 2, 3, 7, 9, 23, 78}
	s2 := []int{5, 6, 8, 45, 67, 70}

	// let it product []int{1,2,3,5,6,7,8,9,45,67,70}

	//naive/brute force approach.

	// it's to loop and compare
	var s3 []int = make([]int, len(s1)+len(s2))

	s3 = append(s1, s2...)

	sort.Ints(s3)

	fmt.Println("first operation", s3)

	//Second step

	// var s4 []int

	var afirst *int
	var bfirst *int

	if len(s1) > 0 {
		afirst = &s1[0]
	}

	if len(s2) > 0 {
		bfirst = &s2[0]
	}

	countFirst := 0
	countSecond := 0
	currentIndex := 0
	secondCompare := 0
	// it's to loop and compare
	var s4 []int = make([]int, len(s1)+len(s2))

	// me := 45;
	// afirst = nil

	for afirst != nil || bfirst != nil {

		if bfirst == nil {
			secondCompare = math.MaxInt32
		} else {
			secondCompare = *bfirst
		}
		if afirst != nil && *afirst < secondCompare {
			// fmt.Println("fool come")
			s4[currentIndex] = *afirst
			countFirst++
			if len(s1) <= countFirst {
				afirst = &s1[countFirst]
			} else {
				afirst = nil
			}

			// fmt.Println("First Did it here", countFirst)
		} else {
			// fmt.Println("fool father")
			s4[currentIndex] = *bfirst
			countSecond++

			if len(s2) <= countSecond {
				bfirst = &s1[countSecond]
			} else {
				bfirst = nil
			}

			// fmt.Println("Second Did it here", countSecond)
		}
		currentIndex++

		cunny()
	}

	fmt.Println("fool", currentIndex, countFirst, countSecond)

	fmt.Println(s3, "Second  Operation")
}

func cunny() {
	fmt.Println("We stand dey")
}
