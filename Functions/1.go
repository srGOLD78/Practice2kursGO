package main

import (
	"fmt"
	"sort"
)

func Find(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func Sorted(slice []string) []string {
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) < len(slice[j])
	})
	return slice
}

func Filter(slice []string, min int) []string {
	var result []string
	for _, elem := range slice {
		if len(elem) >= min {
			result = append(result, elem)
		}
	}
	return result
}

func main() {
	fruits := []string{"apple", "pear", "banana", "orange", "kiwi"}

	fmt.Println("Есть ли банан в строке?", Find(fruits, "banana"))
	fmt.Println("Есть ли манго в строке?", Find(fruits, "mango"))

	sorted := Sorted(fruits)
	fmt.Println("Отсортировано по длине:", sorted)

	filtered := Filter(fruits, 6)
	fmt.Println("Слова длиной >= 6 букв:", filtered)
}
