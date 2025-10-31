package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("введите первый текст: ")
	text1, _ := reader.ReadString('\n')
	fmt.Println("введите второй текст: ")
	text2, _ := reader.ReadString('\n')
	text1 = strings.TrimSpace(text1)
	text2 = strings.TrimSpace(text2)
	words1 := strings.Fields(text1)
	words2 := strings.Fields(text2)

	mapWords1 := make(map[string]int)
	for i := 0; i < len(words1); i++ {
		mapWords1[words1[i]]++
	}
	for i := 0; i < len(words2); i++ {
		if _, ok := mapWords1[words2[i]]; ok {
			fmt.Println(words2[i])
		}
	}
}
