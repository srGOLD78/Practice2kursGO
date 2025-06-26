package main

import "fmt"

func incrementByValue(num int) {
	num++
	fmt.Println("Внутри функции incrementByValue:", num)
}

func incrementByPointer(num *int) {
	*num++
	fmt.Println("Внутри функции incrementByPointer:", *num)
}

func main() {
	value := 7

	fmt.Println("\nПередача по значению:")
	fmt.Println("До вызова:", value)
	incrementByValue(value)
	fmt.Println("После вызова:", value)

	fmt.Println("\nПередача по указателю:")
	fmt.Println("До вызова:", value)
	incrementByPointer(&value)
	fmt.Println("После вызова:", value)
}
