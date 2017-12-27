package main

import "fmt"

func f(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("function call: ", f(1, 41))
	fmt.Println("plusplus: ", plusPlus(1, 40, 1))
}
