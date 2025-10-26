package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is powerful and go is simple"
	text1 := strings.Fields(text)
	words := make(map[string]int)
	for _, k := range text1 {
		words[k]++
	}
	for word, kol := range words {
		fmt.Printf("%s: %d\n", word, kol)
	}
}
