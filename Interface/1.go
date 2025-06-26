package main

import (
	"fmt"
	"math"
)

const pi = 3.1416

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(shape Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
}

func main() {
	rect := Rectangle{Width: 7, Height: 8}
	circle := Circle{Radius: 7.8}

	printArea(rect)
	printArea(circle)
}
