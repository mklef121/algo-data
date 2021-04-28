package main

import (
	"fmt"
	"unicode"
)

func main() {
	//Declares a float with a value of 0.0
	// taxTotal := .0

	fmt.Println(passwordComplexity("aAuuira😅34 ọ  o"))
	fmt.Println(("auuira😅 ọ  o"))
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
