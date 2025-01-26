package main

import "fmt"

/*
Slices have three hidden properties:
	1. length,
	2. a pointer to the hidden array,
	3. and where in the hidden array its starting point is.
When you append to a slice, one or all of these properties get updated.
Which properties get updated depends on whether the hidden array is full or not.
*/

// you can create new slices derived from the contents of arrays and slices. The most common notation is [<low>:<high>].
func main2() {
	long := []int{8, 3, 5, 3, 89, 4, 6, 7, 87, 8}

	fmt.Println(long[2:]) //prints from second index to last
	fmt.Println(long[:4]) //prints from zero(0) index to (4-1) index
	// populate the new slice with values by starting at the low index and then
	// going up to but not including the high index
	fmt.Println(long[1:5]) //prints from index 1 to index (5-1)

	// wllnp array still has reference to the long array
	wllnp := long[:]
	wllnp[0] = 125
	fmt.Println(wllnp, long)

	//this works as if it's a pointer
	// This trick is useful for turning an array into a slice
	kai := editArray(wllnp[:])
	fmt.Println("testing wulolonpo", wllnp, long, kai)

	vee := [4]string{"fg", "gs", "we", "qw"}
	//These two guys(lip and vee) now share the same hidden array because of how slices work
	lip := vee[:] // turns this array to a slice
	lip[0] = "tee"

	var copied []string = make([]string, len(lip))
	copy(copied, lip)
	copied[3] = "I'm copied"
	// A copied slice can still be appended to unlike an array
	copied = append(copied, "Holy-smaoke")
	fmt.Println(vee, lip, cap(vee), cap(lip), "The copied is =>", copied)

	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
	a, b, x := linked()
	fmt.Println("linked arrays function", a, b, x)

	g, s := noLink()
	fmt.Println("No Linked arrays function", g, s)

	q, r, t := capLinked()
	fmt.Println(" capLinked arrays function", q, r, t)
}

func editArray(arr []int) []int {
	arr[3] = 7743

	arr[5] = 54332
	return arr
}

//Using make to Control the Capacity of a Slice
// make(<sliceType>, <length>, <capacity>)
// Due to the complexity of what a slice is and how it works, you can't compare slices to one another.

func genSlices() ([]int, []int, []int) {
	var s1 []int

	// Define a slice using make and set only the length:
	s2 := make([]int, 10)

	// Define a slice that uses both the length and capacity of the slices:
	s3 := make([]int, 10, 50)

	return s1, s2, s3
}

//Controlling Internal Slice Behavior

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	//This forms a new break away
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int, int) {
	// we'll be setting a capacity that's larger than its length:
	// This slice can be expandable till it is 10, once it's length is more than 10, then
	// create another internal array to track it
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1

	s1 = append(s1, 6)
	s1[3] = 99

	s3 := s1
	//create a new array with length of 7 all zero values unless the 6th index that has 45
	x := []int{6: 45}
	//since this superses the cap, it forms a new internal array
	s3 = append(s3, x...)
	// fmt.Println("Shit goes", x, s3)
	s3[3] = 56
	return s1[3], s2[3], s3[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	//Creates an empty array of length 0, then appends s1 to it
	s2 := append([]int{}, s1...)

	//or use this.
	// With the current Go compiler, this is the most memory-efficient way to copy a slice.
	// s4:= append(s1[:0:0], s1...)
	//This uses the seldom-used slice range notation of <slice>[<low>:<high>:<capacity>]
	s1[3] = 99
	return s1[3], s2[3]
}
