package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "привет пока привет Go go это язык программирования влад москва я учу Go всем пока"
	wordCount := make(map[string]int)

	lowerText := strings.ToLower(text)

	words := strings.Fields(lowerText)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("Частота слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
