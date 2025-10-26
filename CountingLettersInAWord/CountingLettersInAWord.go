package main

import "fmt"

func main() {
	text := "gogopher"
	m := make(map[rune]int)
	for _, i := range text {
		m[i]++
	}
	for bukva, element := range m {
		fmt.Printf("%c: %d\n", bukva, element)
	}

}
