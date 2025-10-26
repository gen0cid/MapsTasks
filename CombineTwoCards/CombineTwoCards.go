package main

import "fmt"

func main() {

	map1 := map[string]int{
		"go":  2,
		"is":  3,
		"fun": 1,
	}

	map2 := map[string]int{
		"go":       1,
		"is":       1,
		"powerful": 2,
	}
	map3 := make(map[string]int)

	for i, k := range map1 {
		map3[i] = k
	}

	for j, l := range map2 {
		map3[j] += l
	}
	fmt.Println(map3)
}
