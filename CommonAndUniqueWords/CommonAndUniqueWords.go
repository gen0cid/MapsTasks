package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	//пользователь вводит текс
	fmt.Print("введите первый текст: ")
	text1, _ := reader.ReadString('\n')
	fmt.Print("введите второй текст: ")
	text2, _ := reader.ReadString('\n')

	//удаляю возможные пробелы по краям
	text1 = strings.TrimSpace(text1)
	text2 = strings.TrimSpace(text2)

	//строчку превращаю в массив
	words1 := strings.Fields(text1)
	words2 := strings.Fields(text2)

	//делаю map для двух массивов
	volWords1 := make(map[string]bool)
	volWords2 := make(map[string]bool)

	//перебираю первую сроку и записываю слова в карту
	for _, w := range words1 {
		volWords1[w] = true
	}

	//перебираю вторую строку и записываю слова в карту
	for _, w := range words2 {
		volWords2[w] = true
	}
	//ищу общие слова
	fmt.Println("Общие слова:")
	for _, el := range words2 {
		if volWords1[el] == true {
			fmt.Println(el)
		}
	}
	//ищу уникальные слова первого текста
	fmt.Println("уникальные слова первого текста:")
	for _, el := range words1 {
		if volWords2[el] == false {
			fmt.Println(el)
		}
	}
	//ищу уникальные слова второго текста
	fmt.Println("уникальные слова второго текста:")

	for _, el := range words2 {
		if volWords1[el] == false {
			fmt.Println(el)
		}
	}
}
