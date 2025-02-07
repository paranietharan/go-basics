package main

import (
	"fmt"
	"sync"
	"time"
)

func TestPrintA(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	for i := 0; i <= 100; i++ {
		fmt.Printf("A: %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func TestPrintB(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	for i := 0; i <= 100; i++ {
		fmt.Printf("B: %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// Add two goroutines to the wait group
	wg.Add(2)

	// Start the goroutines
	go TestPrintA(&wg)
	go TestPrintB(&wg)

	// Wait for the goroutines to finish
	wg.Wait()
}
