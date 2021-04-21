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

	add5Value(countTemp)
	fmt.Printf("Original countTemp: %#v\n", countTemp)

	//Turn a  non pointer to a pointer using the "&"
	// The & sign reports the address of the variable
	add5Point(&countTemp)

	fmt.Printf("Second Original countTemp: %#v\n", countTemp)

	a, b := 5, 10

	swap(&a, &b)
	fmt.Println(a == 10, b == 5)

	// add5Point(count1)

	// fmt.Printf(" Original count1: %#v\n", count1)
}

//creating a function that accepts a pointer and the one that does not

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value   :", count)
}

func add5Point(count *int) {
	if count != nil {
		*count += 5
	} else {
		init := 5
		count = &init
		// fmt.Println("Added")
	}

	fmt.Println("add5Point:", *count)
}

func swap(a *int, b *int) {
	// swap the values here
	*a, *b = *b, *a

	// *a = *b
	// *b = *a
}
