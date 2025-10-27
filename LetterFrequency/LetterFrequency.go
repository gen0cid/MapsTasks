package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is easy"
	text1 := strings.ReplaceAll(text, " ", "")
	runes := []rune(text1)
	letters := make(map[rune]int)

	for _, el := range runes {
		letters[el]++
	}

	for key, el := range letters {
		fmt.Printf("%c встречается %d раз\n", key, el)
	}
}
