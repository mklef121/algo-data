package main

import (
	"fmt"
	"time"
)

func printme(x int) {
	fmt.Printf("%d ", x)
}

func main() {
	//Running an anonymous go routine
	go func(x int) {
		fmt.Printf("%d ", x)
	}(10) //pass a parameter to the function

	// This is how you execute a function as a goroutine.
	go printme(15)

	fmt.Println("Only Function peeps")

	// the purpose of the time.Sleep() call is to make the go program wait for it's goroutines to end before exiting
	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}
