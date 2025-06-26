package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Printf("\nРезультаты операций:\n")
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
	fmt.Printf("%.2f %% %.2f = %.2f\n", a, b, float64(int(a)%int(b)))

	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Println("На ноль делить нельзя")
	}

}
