package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) != -1
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"apple", "banana", "orange", "pear", "plum"}

	fmt.Println("Pear index in array: ", Index(strs, "pear"))

	fmt.Println("Does array includes oranges?", Include(strs, "orange"))

	fmt.Println("Is there anything starting with p?", Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println("Does all strings start with p?", All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("Only those that contain e:", Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println("Uppercased: ", Map(strs, strings.ToUpper))
}
