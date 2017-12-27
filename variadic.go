package main

import "fmt"

func sums(nums ...int) int {
	fmt.Println("arguments: ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("sum: ", total)
	return total
}

func main() {
	sums(1, 3, 4)
	sums(13, 15)
	nums := []int{1, 3, 4, 5, 6, 7}

	sums(nums...)
}
