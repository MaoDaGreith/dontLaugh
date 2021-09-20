package main

import (
	"fmt"
	"math"
)

type square struct {
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * 2
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area is", s.area())
}

func main() {
	s1 := square{
		width: 2,
	}

	c1 := circle{
		radius: 12,
	}

	info(s1)
	info(c1)
}
