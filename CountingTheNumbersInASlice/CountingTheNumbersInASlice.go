package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 4, 5, 1, 2}
	numbers := make(map[int]int)
	for _, k := range nums {
		numbers[k]++
	}
	for element, kol := range numbers {
		fmt.Printf("%d: %d\n", element, kol)
	}
}
