package main

import (
	"errors"
	"fmt"
)

func main() {
	var choice int
	var num1, num2, result float64

	i := 1
	for i > 0 {
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")

		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			num1, num2 = getInput()
			result = add(num1, num2)
			fmt.Printf("Result: %.2f\n", result)

		case 2:
			num1, num2 = getInput()
			result = sub(num1, num2)
			fmt.Printf("Result: %.2f\n", result)

		case 3:
			num1, num2 = getInput()
			result = mul(num1, num2)
			fmt.Printf("Result: %.2f\n", result)

		case 4:
			num1, num2 = getInput()
			result, err := div(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %.2f\n", result)
			}

		case 5:
			fmt.Println("Exiting...")
			i = 0 // Exit the loop

		default:
			fmt.Println("Invalid choice! Please select a valid option.")
		}
	}
}

func getInput() (float64, float64) {
	var num1, num2 float64
	fmt.Println("Enter number 1: ")
	fmt.Scan(&num1)
	fmt.Println("Enter number 2: ")
	fmt.Scan(&num2)

	return num1, num2
}

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func sub(num1, num2 float64) float64 {
	return num1 - num2
}

func mul(num1, num2 float64) float64 {
	return num1 * num2
}

func div(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return num1 / num2, nil
}
