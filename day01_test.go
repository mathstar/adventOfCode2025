package main

import "testing"

func TestDay1Part1(t *testing.T) {
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
	assertPart1(day1{}, input, want, t)
}

func TestDay1Part2(t *testing.T) {
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
	assertPart2(day1{}, input, want, t)
}
