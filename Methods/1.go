package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Grades    []int
}

func (s Student) Age() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) AverageGrade() float64 {
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s Student) Status() string {
	avg := s.AverageGrade()
	switch {
	case avg >= 4.5:
		return "отличник"
	case avg >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func main() {
	student := Student{Name: "Волчков Владислав", BirthYear: 2005, Grades: []int{5, 4, 5, 3, 5}}

	fmt.Printf("Студент: %s\n", student.Name)
	fmt.Printf("Возраст: %d лет\n", student.Age())
	fmt.Printf("Средний балл: %.1f\n", student.AverageGrade())
	fmt.Printf("Статус: %s\n", student.Status())
}
