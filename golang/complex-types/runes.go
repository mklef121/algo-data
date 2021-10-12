package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

/*
 * A rune is an alias to the int32 data type. It represents a Unicode code point.
 * A Unicode code point or code position is a numerical value that is usually used to
 * represent a Unicode character. The int32 is big enough to represent the current
* volume of 140,000 unicode characters.
*/
func main() {
	// int32
	a1 := '🦍'
	// int32
	var a2 = 'k'
	// int32
	var a3 rune = '🦡'
	//uint8
	var a4 byte

	fmt.Printf("%c - %s\n", a1, reflect.TypeOf(a1))
	fmt.Printf("%c - %s\n", a2, reflect.TypeOf(a2))
	fmt.Printf("%c - %s\n", a3, reflect.TypeOf(a3))
	fmt.Printf("%c - %s\n", a4, reflect.TypeOf(a4))

	a5 := '🧺'

	//We define two constants; two of them are using escapes.

	//In the first case, the \u is followed by exactly four hexadecimal digits.
	a6 := '\u2665'

	// In the second case, the \U is followed by exactly eight hexadecimal digits.
	a7 := '\U0001F3A8'

	fmt.Printf("%c - %s\n", a5, reflect.TypeOf(a5))
	fmt.Printf("%c - %s\n", a6, reflect.TypeOf(a6))
	fmt.Printf("%c - %s\n", a7, reflect.TypeOf(a7))

	//Go rune Unicode code points

	/*
	 The Unicode code points refer to the characters in the Unicode table.
	 With the %U format verb, we get the Unicode code point.
	 which gives us [U+0066 U+0061 U+006C U+0063 U+006F U+006E] for the string below

	 To get back the string value of one of the unicode codeponts use `string('\u0066')``
	*/

	s1 := "falcon"
	r1 := []rune(s1)
	fmt.Printf("%U\n", r1)

	// Go counting runes
	msg := "one 🐜"
	n1 := len(msg)
	n2 := utf8.RuneCountInString(msg)

	fmt.Println("string length: ", n1) //Outputs 8
	fmt.Println("rune length: ", n2)   //Outputs 5

	//Go runes and bytes
	//A byte in Go is an alias for uint8; it is an "ASCII byte".

	imgGroups := "🐘 🦥 🐋"

	data := []rune(imgGroups)
	fmt.Println(data)

	data2 := []byte(imgGroups)
	fmt.Println(data2)

	//Go loop over runes
	theThing := "one 🐘 and three 🐋"
	// The range keyword generate int32 when used on a string, thus can contain a unicode character

	for idx, e := range theThing {
		fmt.Printf("Char:%s Byte pos: %d \n", string(e), idx)
	}
}
