package main

import (
	"fmt"
)

func f(arg int, source string) int {
	fmt.Println(source)
	return arg + 42
}

func main() {
	f(42, "direct")
	go f(42, "goroutine")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
