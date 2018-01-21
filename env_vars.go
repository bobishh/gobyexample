package main

import (
	"fmt"
	"os"
	s "strings"
)

func main() {
	puts := fmt.Println

	os.Setenv("FOO", "1")
	puts("FOO: ", os.Getenv("FOO"))
	puts("BAR: ", os.Getenv("BAR"))
	puts("------------------------")

	for _, e := range os.Environ() {
		pair := s.Split(e, "=")
		puts(pair)
	}
}
