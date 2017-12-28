package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "result 1"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeouted after 1 second")
	}

	go func() {
		time.Sleep(time.Second)
		ch2 <- "result 2"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(time.Second * 2):
		fmt.Println("timeouted after 2 seconds")
	}
}
