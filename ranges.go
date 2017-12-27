package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum of nums: ", sum)

	for i, num := range nums {
		fmt.Println("i: ", i, "num: ", num)
	}

	kvs := map[string]string{"a": "apple", "b": "orange"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	fmt.Println("only keys:")
	for k := range kvs {
		fmt.Println(k)
	}

	fmt.Println("range over unicode string")
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
