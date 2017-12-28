package main

import "fmt"

func main() {
	queue := make(chan string, 4)
	queue <- "all"
	queue <- "cops"
	queue <- "eat"
	queue <- "mustard"
	close(queue)

	for word := range queue {
		fmt.Println("fetched from queue: ", word)
	}
}
