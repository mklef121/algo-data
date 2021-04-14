package main

import (
	"fmt"
	"time"
)

func main() {
	// Main course meal
	var total float64 = 2 * 13
	fmt.Println("Sub :", total)

	// Drinks
	total = total + (4 * 2.25)
	fmt.Println("Sub :", total)

	// Discount
	total = total - 5
	fmt.Println("Sub :", total)

	// 10% Tip
	tip := total * 0.1
	fmt.Println("Tip  :", tip)

	total = total + tip
	fmt.Println("Total:", total)

	//Then split the bill between two people
	split := total / 2
	fmt.Println("Split: ", split)

	// Reward every 5th visit
	visitCount := 24
	visitCount = visitCount + 1
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward.")
	}

	givenName := "John"
	familyName := "Smith"

	//String concatenation
	fullName := givenName + " " + familyName
	fmt.Println("Hello,", fullName)

	count := 5
	count += 5

	fmt.Println(count)

	count++
	fmt.Println(count)

	count--
	fmt.Println(count)

	count--
	fmt.Println(count)
	name := "Sam"
	name += " Smith"
	fmt.Println("Hello,", name)

	//Zero Values
	var zeroInt int
	var zeroString string

	var emails []string

	fmt.Printf("Emails   : %#v \n The Length is %d \n", emails, len(emails))
	fmt.Printf("The zero value of int is  %+v \n", zeroInt)
	fmt.Printf("The zero value of String is  %#v \n", zeroString)
	//Declare and print a Struct: a type composed of other types
	var startTime time.Time
	fmt.Printf("Start  : %#v \n", startTime)

	emails = []string{"Hello", "Bia"}
	fmt.Printf("Emails   : %#v \n The Length is %d \n", emails, len(emails))

}
