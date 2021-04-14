package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	//Generates a random number between 1 and 5
	random := rand.Intn(5) + 1

	stars := strings.Repeat("*", random)

	fmt.Println(stars)

	medicalActivity()
}

func medicalActivity() {
	//personally add the type so as to evade compile error
	var seed int64 = 1234456789
	rand.Seed(seed)

	var (
		// firstName    string    = "Miracle"
		//Miracle is the Zero value, thus the type is automatically infered by Golang
		firstName              = "Miracle"
		familyName   string    = "Nwabueze"
		peanutAllegy bool      = false
		startUpTime  time.Time = time.Now()
	)
	//Short hand variable declaration. This type must have a zero value
	age := 27
	//Variable names on the left,   variable values on the right..
	Debug, LogLevel, startUpTime := getConfig()
	// myFunc := getConfig;
	// var start, middle, end float32
	// fmt.Println("Shurt up", start, middle, end)
	// 	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	//   fmt.Println(name, left, right, top, bottom)
	fmt.Println(firstName)
	fmt.Println(familyName)
	fmt.Println(age)
	fmt.Println(peanutAllegy)
	fmt.Println(startUpTime)
	fmt.Println(LogLevel)
	fmt.Println(Debug)
}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
