package main

import "fmt"

func vals() (int, int) {
	return 3, 6
}

func main() {
	a, b := vals()
	fmt.Println("values: ", a, b)

	// if you don't care about one of them
	_, c := vals()
	fmt.Println(c)
}
