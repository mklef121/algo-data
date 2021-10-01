package main

import (
	"fmt"
)

/*
A pointer is the memory address of a variable.
You willl need to de-reference a pointer to get it's value
*/

type aStructure struct {
	field1 complex128
	field2 int
}

func main() {
	// use   the * character to get the value in an address
	// Use the & character to get the address of a variable

	a := []string{"me", "uou", "her"}

	fmt.Println(a)

	addData(a)

	fmt.Println(a)

	b := 45.6
	processPointer(&b)

	//B's value is updated as  45.6 * 45.6 and can be seen here even though it is not returned since
	// The pointer is updated inside the function
	fmt.Println(b)

	val := bothPointers(&b)

	fmt.Println(val, *val)

	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)

	// Pointer to f
	fP := &f
	//fP variable is now a variable, having the memory address of the value of variable `f`
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)

	// Check for empty structure
	// Ken is a pointer address to aStructure struct
	//Ken will be nil now since it's just an address with no pointing to a real value
	var ken *aStructure
	var me *int32
	fmt.Println(ken, me)
	// fmt.Printf("%+v\n", fP)
	var pill aStructure
	fmt.Printf("%+v\n", pill)

	if ken == nil {
		ken = new(aStructure)

		fmt.Printf("%+v\n", *ken)
	}
}

func processPointer(x *float64) {
	*x = *x * *x
}

//This is a function that requires a float64 parameter as input and returns a pointer to a float64.
func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func addData(b []string) []string {
	b = append(b, "Hi mama")

	return b
}
