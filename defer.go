package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("/tmp/defer.txt")
	defer closeFile(file)
	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("creating ...")
	file, error := os.Create(path)
	if error != nil {
		panic(error)
	}
	return file
}

func writeFile(f *os.File) {
	fmt.Println("writing ...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing ...")
	f.Close()
}
