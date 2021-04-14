package main

import (
	"fmt"
	"time"
)

/**
 *  A pointer is not a value itself, and you can't do anything useful with a
 * pointer other than getting a value using it. You can think of a pointer as directions
 * to a value you want, and to get to the value, you must follow the directions.
 * If you use a pointer, Go won't make a copy of the value when passing a pointer to a function.
 */
func main() {
	// Varibles are stored in the stack by go. But Pointers are stored in the Heap.
	// Gabbage collector clears it when there is no longer a reference to the pointer

	//pointers created using *var* will always have a value of nill
	var count1 *int

	// While others will already have a value address associated with them.

	// This *new* function is intended to be used to get some memory for a type and return a pointer to that address
	// The deferenced value will be a zero value of it's type
	count2 := new(int)

	countTemp := 5

	// Using &, create a pointer from the existing variable:
	count3 := &countTemp

	// create a pointer from some types without a temporary variable.
	t := &time.Time{}
	var hits bool
	if count3 == nil {
		hits = true
	}

	// fmt.Println(count1, count2, count3)
	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time  : %#v\n", t)
	fmt.Println(hits)

	// Getting a Value from a Pointer

	fmt.Printf("count2: %#v\n", *count2)
	fmt.Printf("count3: %#v\n", *count3)
	fmt.Printf("time : %#v\n", *t)

	fmt.Printf("time  : %#v\n", t.String())

	//this will fail since we are tring to dereference a nil value
	// fmt.Printf("count1: %#v\n", *count1)
	/*
		else use
		if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
		}
	*/

}
