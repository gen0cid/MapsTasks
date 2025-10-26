package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 3, 2, 1, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 5, 4, 3, 2, 1}
	CountingFigure := make(map[int]int)
	for _, k := range nums {
		CountingFigure[k]++
	}

	fmt.Println(CountingFigure)
	for key, s := range CountingFigure {
		if s == 1 {
			fmt.Println("уникальный элемент:", key)
			break
		}
	}
}
