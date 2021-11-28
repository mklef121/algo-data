package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	buffer := []byte("Data to write\n")
	f1, err := os.Create("./f1.txt")

	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}

	defer f1.Close()

	//First formart of writing
	n, err := fmt.Fprintf(f1, string(buffer))

	fmt.Printf("wrote %d bytes\t using fmt.Fprintf \n", n)

	f2, err := os.Create("./f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()

	n, err = f2.WriteString(string(buffer))

	fmt.Printf("wrote %d bytes\t using file.WriteString \n", n)

	f3, err := os.Create("./f3.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	//This function returns a bufio.Writer, which satisfies the io.Writer interface.
	w := bufio.NewWriter(f3)

	n, err = w.WriteString(string(buffer))
	fmt.Printf("wrote %d bytes\t using  bufio.NewWriter \n", n)
	w.Flush()

	f := "./f4.txt"
	f4, err := os.Create(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	for i := 0; i < 5; i++ {
		//Writes to an
		n, err = io.WriteString(f4, string(buffer))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("wrote %d bytes using io.WriteString\n ", n)
	}
	// Append to a file
	// os.OpenFile() provides a better way to create or open a file for writing.
	//os.O_APPEND is saying that if the file already exists, you should append to it instead of truncating it.
	f4, err = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()
	// Write() needs a byte slice
	n, err = f4.Write([]byte("Put some more data at the end.\n"))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes using file.Write\n", n)
}
