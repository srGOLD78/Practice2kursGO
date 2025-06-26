package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	switch operator {
	case "+":
		fmt.Printf("Результат: %.2f\n", a+b)
	case "-":
		fmt.Printf("Результат: %.2f\n", a-b)
	case "*":
		fmt.Printf("Результат: %.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
		} else {
			fmt.Printf("Результат: %.2f\n", a/b)
		}
	default:
		fmt.Println("Ошибка: неизвестный оператор")
	}
}
