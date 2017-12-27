package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextVal := intSeq()
	fmt.Println("next: ", nextVal())
	fmt.Println("next: ", nextVal())
	fmt.Println("next: ", nextVal())
	fmt.Println("next: ", nextVal())
	reseted := intSeq()

	fmt.Println("another: ", reseted())
}
