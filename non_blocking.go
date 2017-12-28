package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	select {
	case msg := <-messages:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("message sent")
	default:
		fmt.Println("message not sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message: ", msg)
	case sig := <-signals:
		fmt.Println("received signal: ", sig)
	default:
		fmt.Println("no activity")
	}
}
