package main

import "fmt"

func send(sendChannel chan<- string, msg string) {
	sendChannel <- msg
}

func rec(recChannel <-chan string) {
	msg := <-recChannel
	fmt.Println("Received : " + msg)
}

func sample(rec <-chan string, send chan<- string) {
	msg := <-rec
	send <- msg
}

func main() {
	ch := make(chan string, 1)

	send(ch, "Hello, Go!")
	rec(ch)

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	send(ch1, "Hi, I am Paranie")
	sample(ch1, ch2)
	rec(ch2)
}
