package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// input := 5
	input := -10
	// fmt.Print(-10 % 2)
	//Implementing the Initial if Statements
	if erro := validate(input); erro != nil {
		fmt.Println(erro, time.Monday)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}

	//Fizbuzz problem
	// FizzBuzz()

	rangeLoop()
	maxCountWord()
}

func validate(input int) error {
	if input < 10 {
		return errors.New("input can't be a negative number")
	} else if input > 100 {
		return errors.New("input can't be over 100")
	} else if input%7 == 0 {
		return errors.New("input can't be divisible by 7")
	} else {
		return nil
	}
}

/*
"FizzBuzz."
The rules are as follows:
• Write a program that prints out the numbers from 1 to 100.
• If the number is a multiple of 3, print "Fizz."
• If the number is a multiple of 5, print "Buzz."
• If the number is a multiple of 3 and 5, print "FizzBuzz."
*/

func FizzBuzz() {

	for i := 1; i <= 100; i++ {
		determineFizStatus(i)
	}
}

func determineFizStatus(num int) {
	if num%3 == 0 && num%5 == 0 {
		fmt.Print(strconv.Itoa(num) + " ")
		fmt.Println("FizzBuzz.")
	} else if num%3 == 0 {
		fmt.Print(strconv.Itoa(num) + " ")
		fmt.Println("Fizz.")
	} else if num%5 == 0 {
		fmt.Print(strconv.Itoa(num) + " ")
		fmt.Println("Buzz.")
	}
}

func rangeLoop() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}

func maxCountWord() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	type Largest struct {
		count int
		name  string
	}
	var maxCountPair Largest
	// maxCountPair = {Number: 2, Text:"fool"}
	for name, count := range words {
		//if it is empty, set the first guy
		if (Largest{}) == maxCountPair {
			maxCountPair = Largest{count, name}
		} else if count > maxCountPair.count {
			maxCountPair = Largest{count, name}
		}
	}

	fmt.Println("Most popular word: ", maxCountPair.name, " \n With a count of  :", maxCountPair.count)
}
