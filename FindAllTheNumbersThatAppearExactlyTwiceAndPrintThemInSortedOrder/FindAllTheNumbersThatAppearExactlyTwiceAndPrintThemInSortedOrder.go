package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 1, 2, 4, 7, 1, 8, 9, 2, 2, 7, 7, 8}
	twice := []int{}
	twiceNumbers := make(map[int]int)
	for _, el := range nums {
		twiceNumbers[el]++
	}
	for key, el := range twiceNumbers {
		if el == 2 {
			twice = append(twice, key)
		}
	}
	sort.Ints(twice)
	fmt.Print("Числа, встречающиеся 2 раза:", twice)
}
