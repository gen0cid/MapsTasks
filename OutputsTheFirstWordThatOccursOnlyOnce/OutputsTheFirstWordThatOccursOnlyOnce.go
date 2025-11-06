package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("введите слово: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	runes := []rune(text)
	letters := make(map[rune]int)
	for _, el := range runes {
		letters[el]++
	}

	duplikate := true

	for _, el := range runes {
		if letters[el] > 1 {
			duplikate = false
			break
		}
	}
	fmt.Println(duplikate)
}
