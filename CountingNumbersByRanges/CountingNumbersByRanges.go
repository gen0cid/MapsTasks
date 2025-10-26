package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 5
	quantity := make(map[string]int)

	for i := 0; i < len(nums); i++ {
		if nums[i] < n {
			quantity["less"]++
		} else if nums[i] > n {
			quantity["greater"]++
		} else {
			quantity["equal"]++
		}
	}
	fmt.Println(quantity)
}
