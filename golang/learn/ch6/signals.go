package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught:", sig)
}

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())

	//We create a channel with data of type os.Signal because all channels must have a type.
	sigs := make(chan os.Signal, 1)

	//We now tell sigs to handle all signals coming through this process
	signal.Notify(sigs)

	//To handle a limited number of signals do so
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGINFO)

	start := time.Now()
	go func() {

		for {
			//Wait until you read data (<-) from the sigs channel and store it in the sig variable.
			sig := <-sigs
			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execution time:", duration)

			case syscall.SIGINFO:
				handleSignal(sig)
				// do not use return here because the goroutine exits // but the time.Sleep() will continue to work!
				os.Exit(0)
			default:
				fmt.Println("Caught:", sig)
			}

		}
	}()

	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}
}
