package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("введите предложение: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	words := strings.Fields(text)
	fmt.Println(words)
	wordsMap := make(map[string]int)

	for _, el := range words {
		wordsMap[el]++
	}
	for _, el := range words {
		if wordsMap[el] == 1 {
			fmt.Println(el)
		}
	}
}
