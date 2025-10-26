package main

import "fmt"

func main() {
	grades := map[string]int{
		"Masha": 5,
		"Gena":  4,
		"Ivan":  3,
		"Oleg":  5,
	}
	bestMark := 0
	bestStudent := ""

	for name, grade := range grades {
		if grade > bestMark {
			bestMark = grade
			bestStudent = name
		}
	}
	fmt.Println("Лучший студент:", bestStudent, bestMark)

}
