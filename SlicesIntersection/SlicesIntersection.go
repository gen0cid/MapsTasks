package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	c := []int{}

	numbers := make(map[int]int)
	for _, i := range a {
		numbers[i]++
	}

	for _, i := range b {
		if _, ok := numbers[i]; ok {
			c = append(c, i)
		}
	}
	fmt.Println(c)
}
