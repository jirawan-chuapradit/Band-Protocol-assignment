package main

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
