package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) perimeter() int {
	return (r.width + r.height) * 2
}

func main() {
	r := rectangle{width: 10, height: 42}

	fmt.Println("area: ", r.area(), ", perimeter: ", r.perimeter())

	rp := &r

	fmt.Println("area: ", rp.area(), ", perimeter: ", rp.perimeter())
}
