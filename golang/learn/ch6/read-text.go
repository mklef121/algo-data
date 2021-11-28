/*
	In this section you will learn how to read plain text files, as well as using the `/dev/random`
	UNIX device, which offers you a way of getting random numbers.
*/

package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	// wordByWord("./readme.md")
	readRandom()
	fmt.Println(string(readSize("./io-interface.go", 12)))
}

//Read random data from unix /dev/random io
func readRandom() {
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var seed int64
	//The endian-ness has to do with the way different computing systems order multiple bytes of information.
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}

//Read a text line by line
func lineByLine(file string) error {
	fd, err := os.Open(file)

	if err != nil {
		return err
	}

	defer fd.Close()

	// First, you create a new reader to the desired file using a call to bufio.NewReader().
	r := bufio.NewReader(fd)

	//Then, you use that reader with bufio.ReadString() in order to read the input file line by line.
	for {
		//ReadString reads until the first occurrence of delim(`\n`) in the input,
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Print(line)
	}

	return nil
}

//Reading a text file wordby word

func wordByWord(file string) error {
	fd, err := os.Open(file)

	if err != nil {
		return err
	}

	defer fd.Close()

	// First, you create a new reader to the desired file using a call to bufio.NewReader().
	r := bufio.NewReader(fd)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}

		reg := regexp.MustCompile("[^\\s]+") //Matches everything that is not a white space
		words := reg.FindAllString(line, -1)

		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}

	return nil
}

func charByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			return err
		}

		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}

//Reading a specific amount of data from a file

func readSize(file string, size int) []byte {
	fd, err := os.Open(file)

	if err != nil {
		fmt.Println("Cannot open file for reading", err)
		return nil
	}

	defer fd.Close()

	buffer := make([]byte, size)

	n, err := fd.Read(buffer)

	if err != nil {
		fmt.Println("Error reading byte data", err)
		return nil
	}

	// io.EOF is a special case and is treated as such
	if err == io.EOF {
		return nil
	}

	//Since the buffer might not be filled to the brim, We pick the section that was successfully written to
	return buffer[0:n]
}
