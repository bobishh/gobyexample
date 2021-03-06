package main

import "fmt"

type person struct {
	nam string
	age int
}

func main() {
	fmt.Println(person{"Bob", 30})
	fmt.Println(person{"Alice", 32})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
