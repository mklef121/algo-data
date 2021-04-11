package main

import (
	"fmt"
	_ "os"
)

func main() {
	// fmt.Printf("This is a sample Go program!")
	// fmt.Println("hello Bithc")
	// fmt.Println("hello Bithc")
	// var me = "Miracle"
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"
	//Prints without a break or space.
	fmt.Print(v1, v2, v3, v4)
	//inserts a line break in your output,
	fmt.Println()
	//Prints each member with a space in between
	fmt.Println(v1, v2, v3, v4)
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")

	//The formatting for the output is predicted here
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)
}
