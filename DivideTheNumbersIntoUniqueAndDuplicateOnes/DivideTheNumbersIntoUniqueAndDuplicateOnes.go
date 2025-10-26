package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 3, 2, 1, 5, 6, 6, 7, 8, 9, 9, 9}
	numbers := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numbers[nums[i]]++
	}
	fmt.Println("уникальные элементы:")
	for key, el := range numbers {
		if el == 1 {
			fmt.Println(key)
		}
	}
	fmt.Print("повторяющиеся элементы:")
	for key, el := range numbers {
		if el != 1 {
			fmt.Printf("%d (%d раз)\n", key, el)
		}
	}
}
