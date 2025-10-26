package main

import "fmt"

func main() {

	nums := []int{1, 1, 2, 3, 3, 3, 4, 4, 4, 4}
	numbers := make(map[int]int)
	//заполяю карту (число : сколько раз встречается)
	for i := 0; i < len(nums); i++ {
		numbers[nums[i]]++
	}
	var maxel, maxkey int

	for key, el := range numbers {
		if el > maxel {
			maxel = el
			maxkey = key
		}
	}
	fmt.Printf("чаще всего встречается: %d (%d раз)\n", maxkey, maxel)
}
