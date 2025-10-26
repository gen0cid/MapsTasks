package main

import "fmt"

func main() {
	c := 0
	f := 0
	nums := []int{1, 2, 3, 4, 5, 6, 2, 2, 4, 5, 6, 7, 7}
	quantity := make(map[int]int)
	for _, value := range nums {
		quantity[value]++
	}
	fmt.Println(quantity)

	for i, k := range quantity {
		if quantity[i] > c {
			f = i
			c = k
		}

	}
	fmt.Println(f, ":", c)
}
