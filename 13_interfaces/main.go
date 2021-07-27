package main

import (
	"fmt"
	"math"
)

// Shape Define an interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{0, 0, 7}
	rectangle := Rectangle{14, 28}

	fmt.Printf("Circle Area %f\n", getArea(circle))
	fmt.Printf("Rectangle Area %f\n", getArea(rectangle))
}
