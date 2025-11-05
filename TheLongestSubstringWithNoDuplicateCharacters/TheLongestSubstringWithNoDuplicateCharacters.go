package main

import "fmt"

func main() {
	text := "abcda"

	runes := []rune(text)      //превратил строку в руны
	last := make(map[rune]int) //сделал карту (ключ -> руна, значение -> индекс последнего появления)
	currentLengh := 0
	maxLength := 0
	left := 0 //левая граница окна
	for i, el := range runes {
		pos, ok := last[el]
		if (ok == true) && (pos >= left) {
			left = pos + 1
		}
		last[el] = i
		currentLengh = i - left + 1
		if currentLengh > maxLength {
			maxLength = currentLengh
		}
	}

	fmt.Println("Максимальная длина подстроки без повторов:", maxLength)

}
