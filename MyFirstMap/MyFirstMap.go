package main

import "fmt"

func main() {
	money := map[string]int{
		"dollars": 1000,
		"euros":   2000,
	}

	for _, value := range money {
		fmt.Println("->", value)
	}
}
