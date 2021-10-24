package main

import (
	"fmt"
	"unicode/utf8"
)

/*
 * As we saw, indexing a string yields its bytes, not its characters: a string is just a bunch of bytes.
 * That means that when we store a character value in a string, we store its byte-at-a-time representation.
 */
func main() {

	//Hexadecimal string represenation
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	//Prints a mercy string
	fmt.Println(sample, "The Byte String")

	//Prints the hexadecimal representation of the string without the "\x"
	//Produces bd b2 3d bc 20 e2 8c 98
	fmt.Printf("% x\n", sample)

	for i := 0; i < len(sample); i++ {

		//Print each number in the hexadecimal
		fmt.Printf("%x ", sample[i])
	}
	//The %q (quoted) verb will escape any non-printable byte sequences in a string so the output is unambiguous.

	//Produces "\xbd\xb2=\xbc ⌘"
	//From here we can see that "\x20" denotes `space` while "\xe2\x8c\x98"  denotes ⌘
	fmt.Printf("\n%q\n", sample)

	stringExplain()

	rangeLoops()
}

func printString() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	//The + prints the unicode rep of unprintable characters
	fmt.Printf("%+q\n", sample)
}

func printByteString() {
	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	//Convert the above to hex to byte slice by converting each to decimal
	var sample = []byte{189, 178, 61, 188, 32, 226, 140, 152}

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	//The + prints the unicode rep of unprintable characters
	fmt.Printf("%+q\n", sample)
}

func stringExplain() {
	//we create a “raw string”,  so it can contain only literal text.
	const placeOfInterest = `⌘`
	fmt.Printf("\n \n")

	fmt.Printf("plain string: ") //outputs ⌘
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("ASCII-only quoted string: ")
	//outputs "\u2318"
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hexadecimal bytes: ")

	//outputs e2 8c 98
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	//This reminds us that the Unicode character value U+2318, the “Place of Interest” symbol ⌘,
	//is represented by the bytes e2 8c 98, and that those bytes are the UTF-8 encoding of
	//the hexadecimal value 2318.
}

// Code points, characters, and runes
/*

 The Unicode standard uses the term “code point” to refer to the item represented by a single value.

 //Example

  the Unicode code point U+0061 is the lower case Latin letter ‘A’: a

  But what about the lower case grave-accented letter ‘A’, à ? That’s a character,
  and it’s also a code point (U+00E0), but it has other representations.

  For example we can use the “combining” grave accent code point, U+0300, and attach it to the
  lower case letter a, U+0061, to create the same character à.

  The concept of character in computing is therefore ambiguous, or at least confusing, so we use it with care.


  “Code point” is a bit of a mouthful, so Go introduces a shorter term for the concept: rune. The term appears
  in the libraries and source code, and means exactly the same as “code point”, with one interesting addition.

  The Go language defines the word rune as an alias for the type int32, so programs
  can be clear when an integer value represents a code point.

  THus Go can represent
  '⌘' as a rune with integer value 0x2318
*/

// Range loops

//there’s really only one way that Go treats UTF-8 specially, and that is when using a for range loop on a string.
// A for range loop, by contrast, decodes one UTF-8-encoded rune on each iteration.
//Each time around the loop, the index of the loop is the starting position of the current rune,
//measured in bytes, and the code point is its value.

func rangeLoops() {
	const nihongo = "日本語"

	fmt.Printf("\n \n for-range loops of unicodes \n\n")
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	fmt.Printf("\n \n using utf-8 library \n\n")

	// Libraries
	//Using the utf8 library to achieve same thing
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
