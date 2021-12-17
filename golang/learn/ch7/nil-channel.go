package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			fmt.Print("New number generated   ", input)
			sum = sum + input
		case <-t.C:
			//We are converting the channel to a nil channel
			c = nil
			fmt.Println("\n\n The sum is ", sum)
			wg.Done()
		}
	}
}

func send(c chan int) {
	for {
		//Once the channel has been set to nill, this loop breaks and returns
		//Because the nill channel blocks. So it stops this go routin from sending data to it
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	rand.Seed(time.Now().Unix())

	wg.Add(1)
	go add(c)
	go send(c)
	wg.Wait()
}
