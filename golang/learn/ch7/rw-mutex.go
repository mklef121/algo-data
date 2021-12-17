package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Recall that internally the sync.Mutex has a format of

type Mutex struct {
	state int32
	sema  uint32
}

then the sync.RWMutex has the format

type RWMutex struct {
	w           Mutex // this is the sync.Mutex
	writerSem   uint32
	readerSem   uint32
	readerCount int32
	readerWait  int32
}

sync.RWMutex is based on sync.Mutex with the necessary additions and improvements.
So, you might ask, how does sync.RWMutex improve sync.Mutex? Although a single function
is allowed to perform write operations with a sync.RWMutex mutex, you can have multiple
readers owning a sync.RWMutex mutex—this means that read operations are usually faster with sync.RWMutex.

The functions that can help you to work with sync.RWMutex are RLock() and RUnlock(),
which are used for locking and unlocking the mutex for reading purposes, respectively
*/

var Password *secret
var wg sync.WaitGroup

type secret struct {
	RWM      sync.RWMutex
	password string
}

//The Change() function makes changes to the shared variable Password and therefore needs to use the Lock()
// function, which can be held by a single writer only.
//Note Only a single function is allowed to perform write operations on sync.RWMutex
func Change(pass string) {
	fmt.Println("Change() function")
	Password.RWM.Lock()
	fmt.Println("Change() Locked")
	time.Sleep(4 * time.Second)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() UnLocked")
}

func show() {
	defer wg.Done()
	Password.RWM.RLock()
	fmt.Println("Show function locked!")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass value:", Password.password)
	defer Password.RWM.RUnlock()
}

func main() {
	Password = &secret{password: "myPass"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("123456")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("54321")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}
