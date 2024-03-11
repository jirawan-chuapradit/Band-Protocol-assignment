package main

import "testing"

func TestBossBabyRevenge_boss_baby_shot_more_than_kids_expected_good_boy(t *testing.T) {
	text := "SRSSRRR"
	want := "Good boy"
	if got := BossBabyRevenge(text); got != want {
		t.Errorf("BossBabyRevenge() = %v, want %v", got, want)
	}
}

func TestBossBabyRevenge_boss_baby_initiate_fight_expected_bad_boy(t *testing.T) {
	text := "RSSRR"
	want := "Bad boy"
	if got := BossBabyRevenge(text); got != want {
		t.Errorf("BossBabyRevenge() = %v, want %v", got, want)
	}
}

func TestBossBabyRevenge_last_shot_no_revenge_expected_bad_boy(t *testing.T) {
	text := "SSSRRRRS"
	want := "Bad boy"
	if got := BossBabyRevenge(text); got != want {
		t.Errorf("BossBabyRevenge() = %v, want %v", got, want)
	}
}

func TestBossBabyRevenge_all_shots_have_been_revenged_expected_good_boy(t *testing.T) {
	text := "SSRR"
	want := "Good boy"
	if got := BossBabyRevenge(text); got != want {
		t.Errorf("BossBabyRevenge() = %v, want %v", got, want)
	}
}

func TestBossBabyRevenge_one_shot_no_revenge_expected_bad_boy(t *testing.T) {
	text := "SRRSSR"
	want := "Bad boy"
	if got := BossBabyRevenge(text); got != want {
		t.Errorf("BossBabyRevenge() = %v, want %v", got, want)
	}
}

func TestBossBabyRevenge_all_shots_have_been_revenged_2_expected_good_boy(t *testing.T) {
	text := "SSRSRR"
	want := "Good boy"
	if got := BossBabyRevenge(text); got != want {
		t.Errorf("BossBabyRevenge() = %v, want %v", got, want)
	}
}
