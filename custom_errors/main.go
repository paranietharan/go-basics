package main

import (
	"errors"
	"fmt"
)

// Custom error struct
type argError struct {
	arg     int
	message string
}

// Implement the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("Error: %d - %s", e.arg, e.message)
}

// Recursive factorial function with error handling
func factorial(num int) (int, error) {
	if num < 0 {
		return -1, &argError{num, "number must be greater than or equal to zero!"}
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
	testValues := []int{5, -3, 0}

	for _, num := range testValues {
		result, err := factorial(num)
		if err != nil {
			var ae *argError
			if errors.As(err, &ae) {
				fmt.Println(ae)
			} else {
				fmt.Println("Unknown error:", err)
			}
		} else {
			fmt.Printf("Factorial of %d is %d\n", num, result)
		}
	}
}
