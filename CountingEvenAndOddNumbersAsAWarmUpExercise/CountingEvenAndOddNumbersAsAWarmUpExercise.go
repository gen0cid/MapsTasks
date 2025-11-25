package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenAndOdd := make(map[string]int)
	for _, k := range nums {
		if k%2 == 0 {
			evenAndOdd["even"]++
		} else {
			evenAndOdd["odd"]++
		}
	}
	fmt.Print(evenAndOdd)
}
