package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("введите строку: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	runes := []rune(text)
	fmt.Println(runes)

	currentLength := 1
	maxLength := 0
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	fmt.Println(maxLength)

}
