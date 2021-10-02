package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100

	startChar := "!"

	fmt.Println("the first", startChar[0], string(startChar[0]), string(89))

	// Each random number generator needs a seed to start producing numbers. The seed is used
	// for initializing the entire process and is extremely important because if you always
	// start with the same seed, you will always get the same sequence of pseudo-random numbers.
	SEED := time.Now().Unix()

	arguments := os.Args
	switch len(arguments) {
	case 2:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		t, err := strconv.Atoi(arguments[1])
		if err != nil {
			MIN = t
			MAX = MIN + 100
		}
	case 3:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		t, err := strconv.Atoi(arguments[1])
		if err == nil {
			MIN = t
		}
		t, err = strconv.Atoi(arguments[2])
		if err == nil {
			MAX = t
		} else {
			MAX = MIN + 100
		}
	case 4:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		t, err := strconv.Atoi(arguments[1])
		if err == nil {
			MIN = t
		}
		t, err = strconv.Atoi(arguments[2])
		if err == nil {
			MAX = t
		} else {
			MAX = MIN + 100
		}
		t, err = strconv.Atoi(arguments[3])
		if err == nil {
			TOTAL = t
		}
	case 5:
		t, err := strconv.Atoi(arguments[1])
		if err == nil {
			MIN = t
		}
		t, err = strconv.Atoi(arguments[2])
		if err == nil {
			MAX = t
		} else {
			MAX = MIN + 100
		}
		t, err = strconv.Atoi(arguments[3])
		if err == nil {
			TOTAL = t
		}
		temp, err := strconv.ParseInt(arguments[4], 10, 64)
		if err == nil {
			SEED = temp
		}
	default:
		fmt.Println("Using default values!")
	}

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		myrand := random(MIN, MAX)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}
