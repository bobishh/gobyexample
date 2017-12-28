package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("start working...")
	time.Sleep(time.Second)
	done <- true
}

func main() {
	done := make(chan bool, 1)
	worker(done)
}
