package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
* What is expected in this file ?
* So we want to match comma seperated input in the following format
* Surname, Firstname, tel number

*
 */
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <surnName, firstName, PhoneNumber> record.")
		return
	}

	s := arguments[1]
	err := matchRecord(s)

	fmt.Println(err)
}

func matchRecord(s string) bool {
	splited := strings.Split(s, ",")

	if len(splited) != 3 {
		fmt.Println("Fields not comma separeated")
		return false
	}

	//surn name check
	if !matchNameSurn(splited[0]) {
		return false
	}

	//first name check
	if !matchNameSurn(splited[1]) {
		return false
	}

	return matchTel(splited[2])
}

func matchTel(s string) bool {
	str := []byte(s)

	reg := regexp.MustCompile(`^\d+$`)

	return reg.Match(str)
}

// Matching names and surnames
func matchNameSurn(s string) bool {
	t := []byte(s)

	//What the regular expression does is match anything that begins with an uppercase letter ([A-Z]) and continues with any number of lowercase letters ([a-z]*).
	reg := regexp.MustCompile(`^[A-Z][a-z]+`)

	return reg.Match(t)
}
