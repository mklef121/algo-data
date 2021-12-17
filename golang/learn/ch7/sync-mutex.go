package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

//The sync.Mutex type is the Go implementation of a mutex.
/*
The interesting work is done by the sync.Lock() and sync.Unlock() functions,
which can lock and unlock a sync.Mutex variable, respectively. Locking a mutex means
that nobody else can lock it until it has been released using the sync.Unlock() function.
*/

//This is the structure of a mutex
type Mutex struct {
	state int32
	sema  uint32
}

var m sync.Mutex
var v1 int

//This function makes changes to the value of v1. The critical section begins here.
func change(i int) {
	//So mutex protects this code block, and would not allow another goroutine work on it until
	// It is unlocked.
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1

	if v1 == 10 {
		v1 = 0
		fmt.Print("* ")
	}

	//This is the end of the critical section. Now, another goroutine can lock the mutex.
	m.Unlock()

}

//This function is used for reading the value of v1—therefore it should use a mutex to make the process concurrently safe
func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	fmt.Printf("%d ", read())

	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}

	wg.Wait()
	fmt.Printf(" final -> %d\n", read())
}
