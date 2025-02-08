package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)

	// channel buffering
	messages := make(chan string, 2)

	messages <- "Hello"
	messages <- "I am Paranie"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
