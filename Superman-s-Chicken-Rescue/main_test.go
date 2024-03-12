package main

import (
	"fmt"
	"testing"
)

func Test_maxProtectedChickensTC1(t *testing.T) {
	roof := 5
	chickens := []int{2, 5, 10, 12, 15}
	wantMaxProtectedChickens := 2
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}

func Test_maxProtectedChickensTC2(t *testing.T) {
	roof := 10
	chickens := []int{1, 11, 30, 34, 35, 37}
	wantMaxProtectedChickens := 4
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}

func Test_maxProtectedChickensTC3(t *testing.T) {
	roof := 5
	chickens := []int{1, 2, 3, 4, 5}
	wantMaxProtectedChickens := 5
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}

func Test_maxProtectedChickensMinRoofAndMinChicken(t *testing.T) {
	roof := 1
	chickens := []int{1}
	wantMaxProtectedChickens := 1
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}

func Test_maxProtectedChickensMinRoofAndMaxChicken(t *testing.T) {
	roof := 1

	// mock chickens
	size := 1_000_000
	chickens := []int{}

	// Populate the slice with numbers increased by 1
	for i := 1; i <= size; i++ {
		chickens = append(chickens, i+1)
	}
	fmt.Println(len(chickens))
	wantMaxProtectedChickens := 1
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}

func Test_maxProtectedChickensNormalRoofAndMaxChicken(t *testing.T) {
	roof := 3

	// mock chickens
	maxLength := 1_000_000_000
	size := 1_000_000
	chickens := []int{}

	// Populate the slice with numbers increased by 1
	for i := 1; i <= maxLength; i++ {
		chickens = append(chickens, i+1)
		if len(chickens) == size {
			break
		}
	}
	wantMaxProtectedChickens := 3
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}

func Test_maxProtectedChickensNormalRoofAndMaxChickenIncreaseByTwo(t *testing.T) {
	roof := 4

	// mock chickens
	maxLength := 1_000_000_000
	size := 1_000_000
	chickens := []int{}

	// Populate the slice with numbers increased by 2
	for i := 1; i <= maxLength; i = i + 2 {
		chickens = append(chickens, i+2)
		if len(chickens) == size {
			break
		}
	}
	// chickens = 1, 3, 5, 7, 9, ...
	wantMaxProtectedChickens := 2
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}

func Test_maxProtectedChickensNormalRoofAndMaxChicken2(t *testing.T) {
	roof := 4

	chickens := []int{1, 3, 5, 7, 9, 999_999_997, 999_999_998, 999_999_999, 1_000_000_000}
	wantMaxProtectedChickens := 4
	if gotMaxProtectedChicken := maxProtectedChickens(roof, chickens); gotMaxProtectedChicken != wantMaxProtectedChickens {
		t.Errorf("maxProtectedChickens() = %v, want %v", gotMaxProtectedChicken, wantMaxProtectedChickens)
	}
}
