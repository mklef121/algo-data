package main

import (
	"fmt"
	"sync"
)

/*
Writing the value val to channel ch is as easy as writing `ch <- val``. The arrow shows the direction of the value.

You can read a single value from a channel named c by executing <-c.
You can save that value into a variable using aVar := <-c.

*/

//This function just writes a value to the channel and immediately closes it.
func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

//This function just sends the true value to a bool channel.
func printer(ch chan bool) {
	ch <- true
}

//This function accepts a channel parameter that is available for writing only.
func printerExplained(ch chan<- bool) {
	ch <- true
}

//The first channel is for reading a value, the second is a send only channel
func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2):", x)
	in <- x
	return
}

func main() {

	//This channel is buffered with a size of 1
	c := make(chan int, 1)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read:", <-c)

	///a technique for determining whether a channel is closed or not
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	waitGroup.Wait()

	//Here, we make an unbuffered channel, and we create five goroutines without
	// any synchronization as we do not use any Add() calls.
	var ch chan bool = make(chan bool)
	for i := 0; i < 5; i++ {
		//this guy will be problematic because multiple go routines are writing to a particular channel
		// This could lead to a race condition
		go printer(ch)
	}

	// Range on channels
	// IMPORTANT: As the channel c is not closed,
	// the range loop does not exit by its own.
	n := 0
	for i := range ch {
		fmt.Println(i, "The Val")
		if i == true {
			n++
		}
		if n > 3 {
			fmt.Println("n:", n)
			close(ch)
			break
		}
	}

	//When trying to read from a closed channel, we get the zero value of its data type,
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
