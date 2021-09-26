package main

import "fmt"

func main() {

	// Byte slice
	b := make([]byte, 12)

	//byte has o zero value of 0
	// An empty byte slice contains zeros—in this case, 12 zeros.
	fmt.Println("Byte slice:", b, "cap: ", cap(b), "length: ", len(b))

	// b now points to a different memory location than before, which is where "Byte slice €" is stored.
	b = []byte("Byte slice €")
	fmt.Println("Byte slice: ", b, "cap: ", cap(b), "length: ", len(b))

	c := []byte{32, 226, 130, 172} // byte representation of `€` unicode

	//convert and print byte slice contents as text
	aStr := string(c)
	fmt.Println(aStr)

}
