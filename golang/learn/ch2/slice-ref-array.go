package main

import "fmt"

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	//we connect S0 with the first element of the array
	var S0 = a[0:1]
	fmt.Println("S0: ", S0)

	//Change the value of the first element
	S0[0] = "S0"

	var S12 = a[1:3]
	fmt.Println("S12 :", S12)

	S12[0] = "S12_0"
	S12[1] = "S12_1"

	// These two changes above will also change the contents of `a`

	fmt.Println("a:", a)

	//This happened because the slices created above uses the same underlying array `a`

	// Changes to slice -> changes to array
	change(S12)
	fmt.Println("change a:", a)

	// capacity of S0 is 4
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	// Adding 4 elements to S0
	// The addition of three more elements makes the length and capacity the same
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1"

	fmt.Println("New S0:", S0, a, S12, cap(S0)) // SO S0 and S12 are still connected to `a`

	//As the capacity of S0 changes, it is no longer connected to the same underlying array (a).
	// Changing the capacity of S0
	// Not the same underlying array anymore!
	S0 = append(S0, "N4")

	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	// This change does not go to S0
	a[0] = "-N1-"
	// This change really affects S12
	// array `a` and slice S12 are still connected because the capacity of S12 has not changed.
	a[1] = "-N2-"

	fmt.Println("S0:", S0)
	fmt.Println("a: ", a)
	fmt.Println("S12:", S12)

}

//This is a function that changes the first element of a slice.
func change(s []string) {
	s[0] = "Change_function"
}
