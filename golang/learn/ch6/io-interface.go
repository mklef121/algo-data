package main

import (
	"bufio"
	"fmt"
	"io"
)

//For the S1 structure, the presented code implements both interfaces in order to
// read user data from the terminal and print data to the terminal, respectively.
type S1 struct {
	F1 int
	F2 string
}
type S2 struct {
	F1   S1
	text []byte
}

// Implementing IO writer and reader interface for S1 struct
func (s *S1) Read(p []byte) (n int, err error) {
	fmt.Print("Give me your name: ")
	fmt.Scanln(&p)
	s.F2 = string(p)
	return len(p), nil
}

func (s *S1) Write(p []byte) (n int, err error) {
	if s.F1 < 0 {
		return -1, nil
	}

	for i := 0; i < s.F1; i++ {
		fmt.Printf("%s ", p)
	}

	fmt.Println()
	return s.F1, nil
}

// implementation of bytes.Buffer.ReadByte
func (s S2) eof() bool {
	return len(s.text) == 0
}
func (s *S2) readByte() byte {
	// this function assumes that eof() check was done before
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

// Implementing io Readable and writable interface for S2 struct

func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return
	}
	l := len(p)

	if l > 0 {
		for n < l {
			p[n] = s.readByte()
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
	}
	return
}

func main() {
	s1var := S1{4, "Hello"}
	fmt.Println(s1var)

	buf := make([]byte, 2)
	_, err := s1var.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Read:", s1var.F2)

	//we call the Write() method for s1var in order to write the contents of a byte slice.
	// writes this to the stdout(terminal)
	_, _ = s1var.Write([]byte("Hello There!"))

	/******     S2     ************/

	s2var := S2{F1: s1var, text: []byte("Hello world!!")}

	// Read s2var.text
	r := bufio.NewReader(&s2var)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("*", err)
			break
		}
		fmt.Println("**", n, string(buf[:n]))
	}
}
