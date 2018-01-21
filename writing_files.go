package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// simple dump to file
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./resources/writes_1", d1, 0644)
	check(err)

	f, err := os.Create("./resources/writes_2")
	check(err)
	// It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	// You can Write byte slices as you’d expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// write strings at once
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// flush changes to storate
	f.Sync()

	// same as with reading - buffered writers from bufio
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// same as f.Sync
	w.Flush()
}
