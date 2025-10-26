package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is powerful and go is simple and go is cool"
	text1 := strings.Fields(text)
	words := make(map[string]int)
	for _, k := range text1 {
		words[k]++
	}

	for element, kol := range words {
		if kol == 1 {
			fmt.Println(element, ":", kol)
		}
	}
}
