package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	jobRole string
	salary  float64
}

func main() {
	e := Employee{
		Person: Person{
			Name: "Paranietharan",
			Age:  23,
		},
		jobRole: "Software Engineer",
		salary:  150.1,
	}

	fmt.Println(e.Name)
	fmt.Println(e.Age)
	fmt.Println(e.jobRole)
	fmt.Println(e.salary)
}
