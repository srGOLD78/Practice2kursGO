package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Меня зовут Влад, я учусь в МТУСИ"

	fmt.Printf("Количество символов: %d\n", len(str))

	substr := "Влад"
	True := strings.Contains(str, substr)
	fmt.Printf("Строка содержит в себе %s? %t\n", substr, True)

	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	fmt.Println("Нижний регистр:", lower)
	fmt.Println("Верхний регистр:", upper)
}
