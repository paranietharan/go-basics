package main

import (
	"errors"
	"fmt"
)

func factorial(num int) (int, error) {
	if num < 0 {
		return -1, errors.New("the number must be positive")
	}

	if num == 0 {
		return 1, nil
	}

	fact, err := factorial(num - 1)
	if err != nil {
		return -1, err
	}
	return num * fact, nil
}

func main() {
	num := -10
	result, err := factorial(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial of", num, "is", result)
	}
}
