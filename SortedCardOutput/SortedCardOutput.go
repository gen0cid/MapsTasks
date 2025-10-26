package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]int{
		"Masha": 5,
		"Gena":  4,
		"Ivan":  3,
		"Oleg":  5,
	}

	names := []string{}

	for i, _ := range grades {
		names = append(names, i)
	}
	sort.Strings(names)
	fmt.Println(names)
	for _, k := range names {
		fmt.Println(k, ":", grades[k])
	}
}
