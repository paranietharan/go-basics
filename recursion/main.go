package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
