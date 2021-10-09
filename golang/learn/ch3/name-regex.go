package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println(`Usage: <utility> string.`)
		return
	}

	surname := arguments[1]

	fmt.Println(matchNameSurn(surname))
}

// Matching names and surnames
func matchNameSurn(s string) bool {
	t := []byte(s)

	//What the regular expression does is match anything that begins with an uppercase letter ([A-Z]) and continues with any number of lowercase letters ([a-z]*).
	reg := regexp.MustCompile(`^[A-Z][a-z]+`)

	return reg.Match(t)
}
