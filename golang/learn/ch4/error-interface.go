package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/*
The error interface looks like this

	type error interface {
		Error() string
	}

thus Any guy that implements that satisfies the interface
*/
type emptyFile struct {
	Ended bool
	Read  int
}

// Implement error interface
func (e emptyFile) Error() string {
	return fmt.Sprintf("Ended with io.EOF (%t) but read (%d) bytes", e.Ended, e.Read)
}

func main() {
	flag.Parse()

	//Run go run error-interface.go /etc/hosts /tmp/doesNotExist /tmp/empty /tmp /tmp/ ./empty.txt
	if len(flag.Args()) == 0 {
		fmt.Println("usage: errorInt <file1> [<file2> ...]")
		return
	}

	for _, file := range flag.Args() {
		err := readFile(file)
		if isFileEmpty(err) {
			fmt.Println("Empty-Error: ", file, err)
		} else if err != nil {
			fmt.Println("Error: ", file, err)
		} else {
			fmt.Println(file, "is OK.")
		}

	}
}

// Check values
func isFileEmpty(e error) bool {
	// Type assertion
	v, ok := e.(emptyFile)

	//This is a type assertion for getting an emptyFile structure from the error variable.
	if ok {
		//if you have read 0 bytes and you are at the end of the file, then it's empty
		if v.Read == 0 && v.Ended == true {
			return true
		}
	}
	return false
}

func readFile(file string) error {

	var err error
	fd, err := os.Open(file)

	if err != nil {
		return err
	}

	defer fd.Close()

	reader := bufio.NewReader(fd)

	n := 0
	for {
		line, err := reader.ReadString('\n')
		n += len(line)

		if err == io.EOF {
			// End of File: nothing more to read
			if n == 0 {
				return emptyFile{true, n}
			}
			break
		} else if err != nil {
			return err
		}

	}

	return err
}
