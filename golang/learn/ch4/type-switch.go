package main

import "fmt"

type Secret struct {
	SecretValue string
}
type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func main() {
	A := Entry{100, "F2", Secret{"myPassword"}}
	Teststruct(make([]int, 0))
	Teststruct(A)
	Teststruct("A string")
	Learn(12)
	Learn('€')
}

func Teststruct(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	case []int:
		fmt.Println("[]int type")
	default:
		fmt.Println(T)
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	//check if this is an int value
	hi, ok := x.(int)
	fmt.Println(hi, ok, "The test")

	// This will PANIC because `anInt` is not `bool` and the `ok` variable is not accessed to test
	// kai := x.(bool)

	// fmt.Println(kai, "The kai factor")

	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}
