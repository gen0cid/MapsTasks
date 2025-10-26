package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is powerful and go is simple and go is cool"
	text1 := strings.Fields(text)
	words := make(map[string]int)
	word := ""
	kol := 0
	for _, k := range text1 {
		words[k]++
	}
	for i, el := range words {
		if el > kol {
			word = i
			kol = el
		}
	}
	fmt.Printf("Самое частое слово: %s (%d раз)\n", word, kol)
}
