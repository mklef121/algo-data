package main

import "fmt"

func main() {
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10)

	fmt.Println("Sum:", sum)
	s := []float64{1.1, 2.12, 3.14}

	sum = addFloats("Adding numbers...", s...)
	fmt.Println("Sum:", sum)
	everything(s)

	str := []string{"One", "Two", "Three"}
	// Cannot directly pass []string as []interface{} // You have to convert it first!
	empty := make([]interface{}, len(str))
	empty = make([]interface{}, len(str))
	for i := range str {
		empty[i] = str[i]
	}

	everything(empty...)

	fmt.Println("\n Defering works \n")

	d1()

	fmt.Println("\n")

	d2()
}

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum = sum + a
	}

	s[0] = -1000
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input)
}

/// The defer keyword

// The defer keyword postpones the execution of a function until the surrounding function returns.

func d1() {
	for i := 3; i > 0; i-- {
		//This is called at the loop instant/
		//3,2,1
		fmt.Println("non-defer looping over: ", i)
		//this will have the updated I value
		//this is called when the d1 function is about returning

		// you get the three values of the i variable of the for loop in reverse order.
		//This is because deferred functions are executed in LIFO order.
		//1,2,3
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		// This is a wrong approach since the anobymous funtion will try getting the value of I on it's own
		//and it gets the last value of I which will be `0`
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

///Better way to implement d2 is in d3

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}
