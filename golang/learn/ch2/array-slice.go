package array_slice

import "fmt"

func main() {

	known_length_array := [4]string{"Zero", "One", "Two", "Three"}
	compiler_length_array := [...]string{"Zero", "One", "Two", "Three"}

	fmt.Println(known_length_array, compiler_length_array)

	// Create an empty slice
	aSlice := []float64{}
	// Both length and capacity are 0 because aSlice is empty
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// Add elements to a slice
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)

	fmt.Println(aSlice, "with length", len(aSlice))

	// A slice with length 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// Now you will need to use append
	t = append(t, -5) //Once a slice has no place left for more elements, you should add new elements to it using append().
	fmt.Println("T :", t)

	// A 2D slice
	// You can have as many dimensions as needed
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	// Visiting all elements of a 2D slice // with a double for loop
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	type me []int

	// Make the initial capacity to be 4.
	make2D := make([][]int, 2, 4)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}

	//Increases the length to 3
	make2D = append(make2D, []int{-0 - 12, -563, -4})
	// increases it's length to 4 and maxes out it's initial capacity set to 4
	// So Go then doubles the capacity each time the length of the slice is about to become bigger than its current capacity
	make2D = append(make2D, []int{-0 - 12, -563, -4})
	fmt.Println("make2D -: ", make2D, len(make2D), cap(make2D))

	theSlice := make([]int, 4, 4)
	addingSlice := []int{-98, -12, -563, -4}
	// Now add four elements
	theSlice = append(theSlice, []int{-1, -2, -3, -4}...)

	theSlice = append(theSlice, addingSlice...)

}
