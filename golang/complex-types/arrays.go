// Go's collection types are array, slice, and map.

package main

import (
	"fmt"
	"os"
)

// DEFINING AN ARRAY
// [<size>]<type>
//For example, [10]int is an array of size 10 that contains ints,
// If your definition didn't have the size, it would not be an array – it'd be a slice
func main() {
	//This size makes it an ARRAY
	array := []int{1, 2, 3, 4}

	//whenever it is created without a definite size, then it's a SLICE
	slice := []int{1, 2, 3, 4, 5, 6}

	// When initializing with data, you can have Go set the size of the array based
	// on the number of elements you initialize it with.

	// For example,
	test := [...]int{9, 9, 9, 9, 9}
	// would create an array of length 5
	//This becaomes an ARRAY, thus the length's set at compile time and is not changeable at runtime.
	// append(test, 46) cannot happen since the length is already fixed

	fmt.Println(array, slice, test)

	//Seeting values at particular indexs with keys
	arr3 := [10]int{1, 9: 10, 4: 5} // means [1 0 0 0 5 0 0 0 0 10]

	//Comparing Arrays

	fmt.Println(arr3)
	fmt.Println(compArrays())
	// comp1, comp2, comp3 := compArrays();

	//Reading from array
	arr := [...]string{"ready", "Get", "Go", "to"}
	fmt.Println(arr[0], arr[3], arr[2])
	second := arr

	//writing to an array
	arr[0] = "Just Ready"
	arr[1] = "To get It"
	//modiications made to the first arr, does not affect the second.
	// No passing arrays by reference
	fmt.Println(arr, second)

	fmt.Print(message())

	findUsers(os.Args[1])

	//Find the longest string passed into the command line

	if longest, shortest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest, "The shortest words was: ", shortest)
	} else {
		fmt.Println("There was an error, No words found while comparing words")
		os.Exit(1)
	}
}

type productPrice struct {
	name       string
	price      float64
	percentage float64
}

func compArrays() (bool, bool, bool) {
	// meshe := map[string]productPrice;
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr5 := [...]int{45, 0, 0, 0, 0}

	// arr7 := []string{}
	// arr8 := []string{}

	// arr7 == arr8;
	// arr4 := [9]int{0, 0, 0, 0, 9}

	// arr1 == arr4 will not work, since the both arrays have different lengths

	//TRUE, TRUE, FALSE
	return arr1 == arr2, arr1 == arr3, arr5 == arr1

	//SLICES cannot be compared using equal to (==). Thus you will have to loop them entirely to compare
}

func message() string {
	m := ""
	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		//Since would not print to console since we are passing it to a variable
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m
}

//SLICES
func getPassedArgs(minArgs int) []string {
	//os.Args holds arguments passed in from the command line
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}

	var args []string

	fmt.Println(os.Args)

	for i := 0; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func findLongest(args []string) (string, string) {
	var longestString string
	var shotestString string

	for i := 0; i < len(args); i++ {
		if len([]rune(args[i])) > len([]rune(longestString)) {
			longestString = args[i]
		}

		if len([]rune(shotestString)) == 0 {
			shotestString = args[i]
		}

		if len([]rune(shotestString)) > len([]rune(args[i])) {
			shotestString = args[i]
		}
	}

	return longestString, shotestString
}

func findUsers(id string) string {
	var users = map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
	mad := []int{3, 45, 6, 8}
	// run := "Hello"...;
	mad = append(mad, mad...)
	mad = append(mad, 6, 8, 9, 4, 3, 2, 3)

	person, found := users[id]

	fmt.Println("The person is ", person, found)
	fmt.Println("Mad array: ", mad)

	return person

}
