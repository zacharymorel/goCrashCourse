package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

// Shared function that the structs that use the interface have
func getArea(s Shape) float64 {
	return s.area()
}

// VALUE POINTER
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// VALUE POINTER
func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	// wrap in Get Area
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

	fmt.Println("Hello World")
}
