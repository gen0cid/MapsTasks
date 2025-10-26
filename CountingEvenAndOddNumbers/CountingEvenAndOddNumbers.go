package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers := make(map[string]int)
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			numbers["odd"] += 1
		} else {
			numbers["even"] += 1
		}
	}
	fmt.Println(numbers)
}
