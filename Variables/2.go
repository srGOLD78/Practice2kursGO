package main

import (
	"fmt"
)

func main() {
	const pi = 3.14
	const e = 2.72

	r := 5.0
	circleArea := pi * (r * r)

	expVal := e * e

	fmt.Println("Значение числа пи:", pi)
	fmt.Println("Значение числа e:", e)
	fmt.Println("Площадь круга", circleArea, "с радиусом", r)
	fmt.Println("e в квадрате равно:", expVal)
}
