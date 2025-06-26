package main

import "fmt"

func main() {
	fruits := []string{"яблоко", "банан", "апельсин", "груша"}
	fmt.Println("Исходный срез:", fruits)
	index := 2

	copy(fruits[index:], fruits[index+1:])
	fruits[len(fruits)-1] = ""
	fruits = fruits[:len(fruits)-1]

	fmt.Println("После удаления:", fruits)
}
