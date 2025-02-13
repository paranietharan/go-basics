package main

import (
	"context"
	"fmt"
	"time"
)

func playGame(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Game over! Time to go home.")
			return
		default:
			fmt.Println("Playing... 🎮")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Mom says: "You have 3 seconds to play!"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go playGame(ctx)

	// Wait for the game to finish
	time.Sleep(5 * time.Second)
	fmt.Println("Back home! 🏠")
}
