/*This code is our package declaration. All Go files must start with one of these.
If you want to run the code directly, you'll need to name it main.
If you don't name it main, then you can use it as a library and import it into other Go code.
When creating an importable package, you can give it any name. All Go files in the same
directory are considered part of the same package, which means all the files must have the same package name.
*/
package main

// Import extra functionality from packages
// In this example, the packages are all from Go's standard library.
// Go's standard library is very high-quality and comprehensive.
// You are strongly recommended to maximize your use of it.
import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// Here, we're declaring a global variable, which is a list of strings, and initializing it with data.
//This type if list is called a slice
var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"‎ایند مالس",
	"Привет, мир",
}

//Here, we're declaring a function. When your code runs, Go automatically calls main to get things started:
func main() {
	fmt.Println(time.Now())
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	// Seed random number generator using the current time
	// Use rand.Seed() before calling any math/rand method, passing an int64 value.
	//Reason is because by default the seed is always the same, the number 1
	//So any time a random number is generated, it will always generated the same numbers
	//https://flaviocopes.com/go-random/
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))
	fmt.Println(rand.Intn(30))

	// Generate a random number in the range of out list
	index := rand.Intn(len(helloList))

	// Call a function and receive multiple return values
	msg, err := hello(index)

	// Handle any errors
	if err != nil {
		log.Fatal(err)
	}
	// Print our message to the console
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		// Create an error, convert the int type to a string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}
