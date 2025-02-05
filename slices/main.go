package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println(nums)
}
