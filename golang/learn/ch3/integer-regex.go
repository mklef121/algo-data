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

	fmt.Println(matchInt(arguments[1]))
}

func matchInt(s string) bool {
	t := []byte(s)

	//No or more sign(+, -) then one or more integers
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}
