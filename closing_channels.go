package main

import "fmt"

func main() {
	jobs := make(chan int, 2)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job: ", j)
			} else {
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 5; i++ {
		jobs <- i
		fmt.Println("sent job: ", i)
	}

	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
	fmt.Println("All done.")
}
