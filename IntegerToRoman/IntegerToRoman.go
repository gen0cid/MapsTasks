package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	lChis := 1

	romanMap := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",

		10: "X",
		20: "XX",
		30: "XXX",
		40: "XL",
		50: "L",
		60: "LX",
		70: "LXX",
		80: "LXXX",
		90: "XC",

		100: "C",
		200: "CC",
		300: "CCC",
		400: "CD",
		500: "D",
		600: "DC",
		700: "DCC",
		800: "DCCC",
		900: "CM",

		1000: "M",
		2000: "MM",
		3000: "MMM",
	}
	fmt.Print("введите число типа int:")
	fmt.Scan(&a)
	c := a
	// нахожу количество цифр в числе "a"
	for {
		if a/10 > 0 {
			lChis++
			a = a / 10
		} else {
			break
		}
	}
	//fmt.Println(3454 / 1000)
	var cifra int
	stepen := 0
	nums := []int{}
	for i := lChis - 1; i >= 0; i-- {
		stepen = int(math.Pow(float64(10), float64(i)))
		cifra = c / stepen
		c = c % stepen
		nums = append(nums, cifra*stepen)
	}
	fmt.Println(nums)

	for _, k := range nums {
		for iterator, element := range romanMap {
			if k == iterator {
				fmt.Print(element)
				break
			}
		}
	}
	fmt.Println()
}
