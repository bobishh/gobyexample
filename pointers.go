package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 42
	fmt.Println("start, i: ", i)

	zeroval(i)
	fmt.Println("zeroed by value, i: ", i)

	zeroptr(&i)
	fmt.Println("zeroed by reference, i: ", i)

	fmt.Println("pointer: ", &i)
}
