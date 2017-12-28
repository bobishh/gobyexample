package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "hello"
	queue <- "world"

	fmt.Println(<-queue)
	fmt.Println(<-queue)
}
