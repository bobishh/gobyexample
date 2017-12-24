package main

import "fmt"

func main() {
	i := 1
	// just like while
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic c-style for
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// no condition, like loop do ... end
	for {
		fmt.Println("loop")
		break
	}

	// continue - skips to next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
