/*
Data Type
int8  ->>> 8-bit signed integer
int16  ->>> 16-bit signed integer
int32  ->>> 32-bit signed integer
int64  ->>> 64-bit signed integer
int	   ->>> 32- or 64-bit signed integer
uint8  ->>> 8-bit unsigned integer
uint16 ->>> 16-bit unsigned integer
uint32 ->>> 32-bit unsigned integer
uint64 ->>> 64-bit unsigned integer
uint   ->>> 32- or 64-bit unsigned integer
float32 ->> 32-bit floating-point number
float64 ->> 64-bit floating-point number
complex64 ->> Complex number with float32 parts
Complex128 ->> Complex number with float64 parts
*/

//The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
package main

import "fmt"

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)
	fmt.Println("Value of c1 and c2: ", c1, c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3)

	cZero := c3 - c3
	fmt.Println("cZero:", cZero)

	x := 12
	k := 5
	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)
	div := x / k //This will always give an integer value even though the division is not perfect
	fmt.Println("div", div)

	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)
	y := 4 / 2.3
	fmt.Println("y:", y)

	//Convert intergers to float64
	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
}
