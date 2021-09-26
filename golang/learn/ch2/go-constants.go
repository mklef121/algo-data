package main

import "fmt"

// Go supports constants, which are variables that cannot change their values
//Strictly speaking, the value of a constant variable is defined at compile time, not
//at runtime—this means that it is included in the binary executable.
//
// Behind the scenes, Go uses Boolean, string, or number as the type for storing constant
//values because this gives Go more flexibility when dealing with constants.
func main() {

	// A Go type is a way of defining a new named type that uses the same underlying type as an existing type.
	type Digit int
	type Power2 float32
	const PI = 3.1415926

	const (
		C1 = "C1C1C1"
		C2 = "C2C2C2"
		C3 = "C3C3C3"
	)

	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	/*
		This constant iota is equivalent to

		const (
			Zero = 0
			One = 1
			Two = 2
			Three = 3
			Four = 4
		)
	*/

	fmt.Println(Three)
	fmt.Println("One:", One)
	fmt.Println(Two)

	// n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".
	// For example, 1 << 5 is "1 times 2, 5 times"
	const (
		p2_0 Power2 = 1 << iota
		_           //the underscore character in a const block with a constant generator iota, which allows you to skip unwanted values

		//iota has the value of 2 and p2_2 is defined as the result of the expression 1 << 2,
		// which is 00000100 in binary representation.
		p2_2 //iota is 2 here so p2_2 = 1 << 2 (1*2)^2
		_
		p2_4 //iota is 4 here so p2_4 = 1 << 4 (1*2)^4
		_
		p2_6
	)

	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)

	twoDEmptySlice := make([][]int, 2)
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Slices Shit:", twoDEmptySlice, twoD)
}
