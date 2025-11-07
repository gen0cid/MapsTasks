package main

import "fmt"

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3
	left := 0
	currentSum := 0
	maxSum := 0

	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum = currentSum

	for right := k; right < len(nums); right++ {
		currentSum += nums[right]
		currentSum -= nums[left]
		left++
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	fmt.Println(maxSum)
}
