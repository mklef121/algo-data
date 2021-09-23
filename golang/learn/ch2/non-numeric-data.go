package nonNumeric

import "fmt"

// Go has support for Strings, Characters, Runes, Dates, and Times.
// However, Go does not have a dedicated char data type
func main() {

	// A Go string is just a collection of bytes and can be accessed as a whole or as an array.
	// A single byte can store any ASCII character—however,
	//  multiple bytes are usually needed for storing a single Unicode character.
	var first int = 20
	// var runeMe = [32,240,159,152,131]
	fmt.Println(string(first), []byte("8 A String 😃"))

	aString := "Hello World! €"
	fmt.Println("First character", string(aString[0]))

	// Runes
	// A rune
	r := '😃'
	fmt.Println("As an int32 value:", r)
	// Convert Runes to text
	// The %c control string in fmt.Printf() prints a rune as a character.
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// Print an existing string as runes
	for _, v := range aString {
		// fmt.Println(v)
		fmt.Printf("%x ", v)
	}
	fmt.Println()
}
