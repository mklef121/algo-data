package main

import (
	"fmt"
	"time"
)

/**
There are times that goroutines take more time than expected to finish—in such situations,
we want to time out the goroutines so that we can unblock the program.
There are two techniques to do such
**/
func main() {

	// First method to timeout a goroutine
	c1 := make(chan string)
	go func() {
		//We are asuming that this function makes a call that takes 3 seconds to complete
		time.Sleep(3 * time.Second)
		fmt.Println("About sending message to channel 1")
		c1 <- "c1 OK"
		fmt.Println("Sent message to channel 1")
	}()

	// res := <-c1
	// fmt.Println("first detect", res)
	select {
	case res := <-c1:
		fmt.Println(res)
	//The purpose of the time.After() call is to wait for the desired time before being executed
	case <-time.After(time.Second):
		fmt.Println("timeout c1")
	}

	//At this point, the subscription to messages from channel one has been put off, so when the go routine notices
	//That the channel has no more subscription, it goes off or times out

	c2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("About sending message to channel 2")
		c2 <- "c2 OK"
		fmt.Println("Sent message to channel 2")
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout c2")
	}

	//At the time a listener is set to time out at 4 seconds, the channel would have sent it's message
	//So this second go routine does not time out
}
