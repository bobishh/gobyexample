package main

import (
	"fmt"
	"strconv"
)

func main() {
	puts := fmt.Println
	f, _ := strconv.ParseFloat("1.234", 64)
	puts(f)

	s := strconv.ParseInt("123", 0, 64)
	puts(s)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) // 0 means infer from string
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135") // convenient way of base-10 parsing
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
