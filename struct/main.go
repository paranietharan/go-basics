package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name  string
	age   int
	email string
}

func newPerson(name string, age int, email string) *Person {
	p := Person{
		name:  name,
		age:   age,
		email: email,
	}

	return &p
}

func main() {
	m := newPerson("Mayu", 26, "Mayu@gmail.com")
	fmt.Println(m)
	fmt.Println(*m)

	fmt.Println("Name : " + m.name)
	fmt.Println("Age : " + strconv.Itoa(m.age))
	fmt.Println("Email Address : " + m.email)

	p := Person{
		name:  "Paranie",
		age:   24,
		email: "Paranietharan20@gmail.com",
	}
	fmt.Printf("\nName: %s\nAge : %d\nEmail: %s\n", p.name, p.age, p.email)
}
