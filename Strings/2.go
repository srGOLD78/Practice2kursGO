package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Меня зовут Влад и я учусь в МТУСИ"

	words := strings.Fields(sentence)

	fmt.Println("Слова в предложении:")
	fmt.Println(words[0])
	fmt.Println(words[1])
	fmt.Println(words[2])
	fmt.Println(words[3])
	fmt.Println(words[4])
	fmt.Println(words[5])
	fmt.Println(words[6])
	fmt.Println(words[7])
}
