package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("here we go: ", i)
	}

	// For loop used as while loop
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value: ", v)
	}

	getName()
}

// Getting user input
/**
The fmt.Scanln() function can help you read user input while the program is already
running and store it to a string variable.
**/

func getName() {
	// Get User Input
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
