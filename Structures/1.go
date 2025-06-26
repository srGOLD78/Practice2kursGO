package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age, course int, avgGrade float64) Student {
	return Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AvgGrade)
}

func (s *Student) Up() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func main() {
	vlad := NewStudent("Волчков Владсилав", 20, 2, 4.2)

	fmt.Println("Информация о студенте:")
	vlad.PrintInfo()

	vlad.Up()

	fmt.Println("\nОбновленная информация:")
	vlad.PrintInfo()
}
