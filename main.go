package main

import "strings"

// S shots taken by Boss Baby,
// R shots taken by kids
// and then compares them to determine whether Boss Baby is a "Good boy" or a "Bad boy"
func BossBabyRevenge(shot string) string {
	kidShot := 0
	shot = strings.ToUpper(shot)

	// boss_baby_initiate_fight
	if strings.HasPrefix(shot, "R") {
		return "Bad boy"
	}

	for _, shot := range shot {
		if shot == 'S' {
			kidShot++
		} else if shot == 'R' && kidShot > 0 {
			kidShot--
		}
	}

	if kidShot <= 0 {
		return "Good boy"
	} else {
		return "Bad boy"
	}
}
