package main

import (
	"fmt"
)

func main() {
	arr1 := [10]int{11, 2, 7, 5, 29, 74, 32, 56, 101, 0}

	for i, v := range arr1 {
		fmt.Printf("index %d, value %d\n", i, v)
	}

	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	var fruits [3]string = [3]string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, fruit: %s\n", index, fruit)
	}
}
