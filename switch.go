package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	// good ol' switch
	switch i {
	case 1:
		fmt.Println(i)
	case 2:
		fmt.Println(i)
	}

	now := time.Now()

	// default case and match with a list of predicates
	switch now.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Happy weekends!")
	default:
		fmt.Println("Go work! It's only " + now.Weekday().String())
	}

	// no expression - if analogue
	switch {
	case now.Hour() < 12:
		fmt.Println("It's early!")
	default:
		fmt.Println("It's afternoon")
	}

	// -- super duper type checker
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
