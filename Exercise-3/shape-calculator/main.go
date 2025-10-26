package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height

}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func getPerimeter(s Shape) {
	switch shape := s.(type) {
	case Circle:
		perimeter := 2 * math.Pi * shape.Radius
		fmt.Println("Perimeter of the circle is :", perimeter)
	case Rectangle:
		perimeter := 2 * (shape.Height + shape.Width)
		fmt.Println("Rectangle:", perimeter)
	case Triangle:
		fmt.Println("Oops i can not find perimeter with just the base and height")
	default:
		fmt.Println("Unknown shape")
	}
}
func describeShape(s Shape) {
	switch shape := s.(type) {
	case Circle:
		fmt.Println("Circle with radius:", shape.Radius)
	case Rectangle:
		fmt.Println("Rectangle:", shape.Width, "x", shape.Height)
	case Triangle:
		fmt.Println("Triangle with base:", shape.Base, "and height:", shape.Height)
	default:
		fmt.Println("Unknown shape")
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{Base: 3, Height: 4},
	}

	for _, shape := range shapes {
		describeShape(shape)
		printArea(shape)
		getPerimeter(shape)
	}
}
