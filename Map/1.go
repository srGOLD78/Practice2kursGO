package main

import "fmt"

func Add(grades map[string]int, name string, grade int) {
	grades[name] = grade
	fmt.Printf("Добавлена оценка для %s: %d\n", name, grade)
}

func Find(grades map[string]int, name string) {
	if grade, exists := grades[name]; exists {
		fmt.Printf("Оценка %s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func Delete(grades map[string]int, name string) {
	if _, exists := grades[name]; exists {
		delete(grades, name)
		fmt.Printf("Оценка для %s удалена\n", name)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func main() {
	grades := make(map[string]int)

	Add(grades, "Алексей", 5)
	Add(grades, "Мария", 4)
	Add(grades, "Иван", 3)

	Find(grades, "Мария")
	Find(grades, "Петр")

	Delete(grades, "Иван")
	Delete(grades, "Петр")

	fmt.Println("\nТекущие оценки:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
}
