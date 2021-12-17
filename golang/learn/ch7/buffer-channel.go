package main

import (
	"fmt"
)

func main() {
	//creating a buffered channel

	//The numbers channel cannot store more than five integers—this is a buffer channel with a capacity of 5.
	numbers := make(chan int, 5)

	counter := 10

	for i := 0; i < counter; i++ {
		//We begin putting data into numbers—however, when the channel is full, it is not going to
		// store more data and the default branch is going to be executed.
		select {
		// This is where the processing takes place
		case numbers <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Print("No space for ", i, " ")
		}
	}

	fmt.Println()
	for {
		select {
		case num := <-numbers:
			fmt.Print("*", num, " ")
		default:
			fmt.Println("Nothing left to read!")
			return
		}
	}
}
