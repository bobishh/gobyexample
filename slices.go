package main

import "fmt"

func main() {
	// slices
	slice := make([]string, 3)
	fmt.Println("empty: ", slice)

	slice[0] = "punk's"
	slice[1] = "not"
	slice[2] = "dead"

	fmt.Println("filled: ", slice)
	fmt.Println("element: ", slice[2])

	fmt.Println("length: ", len(slice))

	slice = append(slice, "but")
	slice = append(slice, "stinks", "a lot")

	fmt.Println("appended: ", slice)

	// creating new slice as an old one's copy
	newSlice := make([]string, len(slice))

	copy(newSlice, slice)

	fmt.Println("sliced from middle: ", newSlice[2:5])
	fmt.Println("sliced from start to final: ", newSlice[:5])
	fmt.Println("sliced from initial to end: ", newSlice[2:])

	declaredAndInitialied := []string{"G", "B", "H"}
	fmt.Println("declaredAndInitialied: ", declaredAndInitialied)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
