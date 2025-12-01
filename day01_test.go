package main

import "testing"

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	want := "3"
	got := day1{}.part1(input)
	if got != want {
		t.Errorf("part1() = %q, want %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	want := "6"
	got := day1{}.part2(input)
	if got != want {
		t.Errorf("part2() = %q, want %q", got, want)
	}
}
