package main

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"unicode"
)

func main() {
	//Declares a float with a value of 0.0
	// taxTotal := .0

	fmt.Println(passwordComplexity("aAuuira😅34 ọ  o"))
	fmt.Println(("auuira😅 ọ  o"))

	//Know heap memory used by program
	spaceOccupiedByNumber()

	goTexts()
}

func passwordComplexity(password string) bool {
	/*	• Have a lowercase letter
		• Have an uppercase letter
		• Have a number
		• Have a symbol
		• Be 8 or more characters long
	*/
	// Convert the password string into rune, which is a type that is safe for multi-byte (UTF-8) characters:
	pwR := []rune(password)
	// Same as pwR := []int32(password)
	var (
		hasUpper  = false
		hasLower  = false
		hasNumber = false
		hasSymbol = false
	)
	// pwR := make(map[string]string, 4)

	if len(pwR) < 8 { //satisfies condition 5
		return false
	}

	for _, value := range pwR {

		// hasUpper = unicode.IsUpper(value) ? true: false

		if unicode.IsUpper(value) {
			hasUpper = true
		}

		if unicode.IsLower(value) {
			hasLower = true

		}

		if unicode.IsPunct(value) || unicode.IsSymbol(value) {
			hasSymbol = true
		}

		if unicode.IsNumber(value) {
			hasNumber = true
		}

	}

	fmt.Println(hasUpper, hasLower, hasNumber, hasSymbol)

	return hasUpper && hasLower && hasNumber && hasSymbol
}

//Numbers

func spaceOccupiedByNumber() {
	//When running into memory issues, It's good We use integers that occupy less space.

	//When using this TotalAlloc (Heap) = 54 MiB
	// var list []int

	//When using this TotalAlloc (Heap) = 54 MiB
	var list []int8

	//Make additional 10 million additions to the List to see how much Heap memory it consumes
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}

	var memo runtime.MemStats

	runtime.ReadMemStats(&memo)

	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", memo.TotalAlloc/1024/1024)

	var a int = 100
	var b float32 = 100
	var c float64 = 100
	fmt.Println(a / 3) //33
	fmt.Println(b / 3) //33.333 332 -> 16bit to contain 33, 16bit to contain  333332: the 32 bit shared
	fmt.Println(c / 3) //33.333333333333336

	intA := math.MaxInt64 //9223372036854775807
	intA = intA + 1       //-9223372036854775808 Adding this takes it to the smallest, what golang calls Wraparound
	fmt.Println("Warapped round Biggest is", intA)
	// Now we'll create a big int. This is a custom type and is not based on Go's int type.
	// We'll also initialize it with Go's highest possible number value:
	bigA := big.NewInt(math.MaxInt64) //9223372036854775807
	bigA.Add(bigA, big.NewInt(1))     //9223372036854775808 But using the big Int does not Wraparound the number

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int   :", intA)
	fmt.Println("Big Int : ", bigA.String())
}

// **Floating Point**
// Go has two floating-point number types, float32 and float64.

//  **Byte**
// The byte type in Go is just an alias for uint8
//This is a number that has 8 bits of storage
// 8 bits have 256 possible combinations of "off" and "on," uint8 has 256 possible integer values from 0 to 255.

func goTexts() {
	comment1 := `This is the BEST
	thing ever!` //This is a raw string. Go inteprets every space and tab as the equivalent value
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!" //Interpreted string
	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	username := "Sir_King_Über"

	// fmt.Println("The length of a special string is", len("😅"), len([]rune("😅"))) // gives us two(2)
	for i := 0; i < len(username); i++ {
		//This way we are accessing the byte (uint8)
		// Reason why len("Ü") will return a length of 2
		// It was encoded using more than one byte
		byteData := username[i]
		fmt.Print("Byte Value: ", byteData, " The string Value: ", string(byteData), " \n")

	}

	fmt.Print("\n")
	//To safely work with interindividual characters of a multi-byte string,
	// you first must convert the strings slice of byte types to a slice of rune types.

	rune_type := []rune(username)

	for i := 0; i < len(rune_type); i++ {
		byteData := rune_type[i]
		fmt.Print("Rune Value: ", byteData, " The string Value: ", string(byteData), " \n")
	}

	//for _,val := range username { //gives us the rune immedaitely
	// }

	// fmt.Println(username[:10], "10th placed element", []rune(username)[:10])

	fmt.Print("\n")

	// var string_type = "🇵🇬"
	// int_values := []uint8(string_type)

	// for _, val := range int_values {
	// 	println("flag number ", val)
	// }

}

// **Rune**
// A rune is a type with enough storage to store a single UTF-8 multi-byte character.
// Remember rune == int32
// The string type itself is not limited to UTF-8 as Go needs to also support text encoding types other than UTF-8
// string not being limited to UTF-8 means there is often an extra step you need to take when working
// with your strings to prevent bugs
// Legacy standards use one byte to encode a single character. UTF-8 uses up to four bytes to encode a single character.
// Go stores all strings as a byte collection

/************* Performing Operations on Strings ******/

// To be able to safely perform operations with text of any kind of encoding,
// single- or multi-byte, it should be converted from a byte collection to a rune collection.
