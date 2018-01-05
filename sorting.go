package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"c", "a", "b"}
	sort.Strings(strings) // sorting is mutable
	fmt.Println(strings)

	ints := []int{12, 3, 5, 1}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)

	fmt.Println("sorted?:", s)
}
