package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	group := make(map[int][]float64)

	for i, _ := range arr {
		a := int(arr[i]/10) * 10
		group[a] = append(group[a], arr[i])
	}
	fmt.Println(group)
}
