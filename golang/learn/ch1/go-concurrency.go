package main

import (
	"fmt"
	"time"
)

/**
The Go concurrency model is implemented using goroutines and channels.
A goroutine is the smallest executable Go entity.

A channel in Go is a mechanism that, among other things, allows goroutines to communicate and exchange data.

The Go scheduler is responsible for the execution of goroutines just like the OS scheduler
is responsible for the execution of the OS threads.
**/
func main() {

	for i := 0; i < 5; i++ {
		go myPrint(i, 5)
	}

	fmt.Println("main for loop is outside")
	time.Sleep(time.Second)
}

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ", "the ", start)
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}
