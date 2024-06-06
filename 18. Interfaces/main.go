package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return math.Pi * 2 * c.radius
}

type shape interface {
	area() float64
	circumference() float64
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f\n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f\n", s, s.circumference())
}

func main() {
	shapes := []shape{
		square{length: 5},
		square{length: 10},
		circle{radius: 2},
		circle{radius: 4},
	}

	for _, v := range shapes {
		printShapeInfo(v)
	}
}
