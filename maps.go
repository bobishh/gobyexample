package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["tomatoes"] = 42
	m["bananas"] = 24

	fmt.Println(m)

	value := m["tomatoes"]
	fmt.Println("printing value: ", value)
	fmt.Println("map length: ", len(m))

	delete(m, "bananas")
	fmt.Println("after deletion", m)

	_, present := m["k2"]
	fmt.Println("present?:", present)

	declaredAndInitialized := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", declaredAndInitialized)
}
