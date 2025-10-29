package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text1 := strings.Fields(text)
	words := make(map[string]int)

	for _, k := range text1 {
		words[k]++
	}

	fmt.Println("\nЧастота слов:")

	for key, el := range words {
		fmt.Printf("%s: %d\n", key, el)
	}

	fmt.Println("\nЧастота букв:")

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
