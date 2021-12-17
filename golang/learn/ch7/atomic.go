package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	val int64
}

//This is a helper function that returns the current value of an int64 atomic variable using atomic.LoadInt64().
func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}
func main() {
	X := 100
	Y := 4
	var waitGroup sync.WaitGroup
	counter := atomCounter{}
	for i := 0; i < X; i++ {
		waitGroup.Add(1)
		go func(no int) {
			//We are creating lots of goroutines that change the shared variable
			defer waitGroup.Done()
			for i := 0; i < Y; i++ {
				//The atomic.AddInt64() function changes the value of the val field of the counter structure variable in a safe way.
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}

	waitGroup.Wait()
	fmt.Println(counter.Value())

}
