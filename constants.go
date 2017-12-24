package main

import (
	"fmt"
	"math"
)

const c string = "constant"

func main() {
	fmt.Println(c)
	const n = 500000000
	const d = 3e20 / n

	fmt.Println("n: ", n)
	fmt.Println("d: ", d)

	fmt.Println("d, casted: ", int64(d))
	fmt.Println("sin n: ", math.Sin(n))
}
