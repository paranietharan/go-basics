package main

import "fmt"

func main() {
	ageMap := make(map[string]int)

	ageMap["Paranie"] = 24
	ageMap["Mayuravel"] = 25

	fmt.Println(ageMap)
	fmt.Printf("Paranie age %d\n", ageMap["Paranie"])
}
