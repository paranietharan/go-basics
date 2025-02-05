package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type rectangele struct {
	width, height float64
}

func (r rectangele) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := rectangele{
		width:  20.0,
		height: 23.5,
	}

	c := circle{
		radius: 10.1,
	}

	fmt.Println(r.area())
	fmt.Println(c.area())
}
