package main

import (
	"fmt"
)

func maxProtectedChickens(roof int, chickens []int) int {
	prevChickenProtected := 0
	for _, min := range chickens {
		max := min + roof

		// Count occurrences of numbers within the range [min, max]
		count := make(map[int]int)
		for _, n := range chickens {
			if n >= min && n < max {
				count[n]++
			} else if n >= max {
				break
			} else if len(count) == roof {
				return len(count)
			}
		}

		if len(count) > prevChickenProtected {
			prevChickenProtected = len(count)
		}
	}

	return prevChickenProtected
}

func main() {
	chickenProtected := maxProtectedChickens(5, []int{2, 5, 10, 12, 15})
	fmt.Println("chicken protected:", chickenProtected)

	chickenProtected = maxProtectedChickens(10, []int{1, 11, 30, 34, 35, 37})
	fmt.Println("chicken protected:", chickenProtected)

}
