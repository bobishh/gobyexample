package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for {
			_, more := <-ticker.C
			if more {
				fmt.Println("TICK-TOCK")
			} else {
				return
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped.")
}
