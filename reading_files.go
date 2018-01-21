package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	puts := fmt.Println
	// read all file from memory
	dat, err := ioutil.ReadFile("./resources/reads")
	check(err)
	puts(string(dat))

	// open file and not loading all contents
	f, err := os.Open("./resources/reads")
	check(err)

	// read 5 bytes from the beginning
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// seek to known location and read from there
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// using io package for the task above
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// seek to 0, 0 is a rewind
	_, err = f.Seek(0, 0)
	check(err)

	// bufio package implements a buffered reader that may be useful
	// both for its efficiency with many small reads and because of
	// the additional reading methods it provides

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// close the file

	f.Close()
}
