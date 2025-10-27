package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is easy"
	text1 := strings.Fields(text)
	words := make(map[string]int)

	for _, k := range text1 {
		words[k]++
	}

	fmt.Println("Частота слов:")

	for key, el := range words {
		fmt.Printf("%s: %d\n", key, el)
	}

	fmt.Println("Частота букв:")

	letters := make(map[rune]int)
	text2 := strings.ReplaceAll(text, " ", "")
	runes := []rune(text2)

	for _, k := range runes {
		letters[k]++
	}

	for key, el := range letters {
		fmt.Printf("%c: %d\n", key, el)
	}
}
