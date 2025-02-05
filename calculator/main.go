package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Sum:", add(10, 5))
	fmt.Println("Difference:", sub(10, 5))
	fmt.Println("Product:", mul(10, 5))

	result, err := div(10, 0) // Testing division by zero
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", result)
	}
}

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func sub(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func mul(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func div(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("division by zero is not allowed")
	}

	return num1 / num2, nil
}
