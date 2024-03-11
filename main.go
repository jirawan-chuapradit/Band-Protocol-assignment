package main

import "strings"

// S shots taken by Boss Baby,
// R shots taken by kids
// and then compares them to determine whether Boss Baby is a "Good boy" or a "Bad boy"
func BossBabyRevenge(shot string) string {
	bossRevenge := 0
	shot = strings.ToUpper(shot)

	if initiateFight := isBossBabyInitiateFight(shot); initiateFight {
		return "Bad boy"
	}

	for _, shot := range shot {
		if shot == 'S' {
			bossRevenge++
		} else if shot == 'R' && bossRevenge > 0 {
			bossRevenge--
		}
	}

	if bossRevenge > 0 {
		return "Bad boy"
	}
	return "Good boy"
}

// check Boss Baby initiate shots to the neighborhood kids
func isBossBabyInitiateFight(shot string) bool {
	return strings.HasPrefix(shot, "R")
}
