package main

import "fmt"

func main() {
	var fruits []string

	fruits = append(fruits, "яблоко")
	fruits = append(fruits, "банан")
	fruits = append(fruits, "апельсин")

	fmt.Println("Содержимое среза:")
	fmt.Println(fruits)
	fruits = append(fruits, "груша", "персик")

	fmt.Println("Добавили еще фруктов")
	fmt.Println(fruits)
}
