package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// You can convert an integer value into a string in two main ways: using string()
// and using a function from the strconv package
func main() {
	// The string() function converts an integer value into a Unicode code point, which is a single character

	// functions such as strconv.FormatInt() and strconv.Itoa() convert an integer
	// value into a string value with the same representation and the same number of characters.

	if len(os.Args) == 1 {
		fmt.Println("Print provide an integer.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Using strconv.Itoa()
	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)

	// Using strconv.FormatInt
	input = strconv.FormatInt(int64(n), 10) //represents numbers the same way as string.
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)

	// Using string()
	input = string(n) //Converts numbers into their unicode representation
	fmt.Printf("string() %s of type %T\n", input, input)

	// The unicode package
	//The unicode standard Go package contains various handy functions for working with Unicode code points
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	fmt.Println("\x77\x00ab\x50", "Im the shit")
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}

	// /The strings package
	// The strings standard Go package allows you to manipulate UTF-8 strings in Go
}
