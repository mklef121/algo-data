package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Println("argument " + arguments[i] + " cannot be converted to float")
			continue
		}

		if i == 1 {
			min = n
			max = n
			continue
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	// Using error variables to differentiate between input types

	var total, nInts, nFloats int
	invalid := make([]string, 0)

	for _, k := range arguments[1:] {
		// Is it an integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		// Is it a float
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		// Then it is invalid
		invalid = append(invalid, k)

	}

	fmt.Println("#Ints: ", nInts, " #FLoats: ", nFloats, "Failed: ", invalid)
}
