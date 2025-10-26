package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is easy"

	WordCount := make(map[string]int)

	words := strings.Fields(text)

	for _, value := range words {
		WordCount[value]++
	}
	fmt.Println(WordCount)
}
