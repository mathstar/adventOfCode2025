package main

import "testing"

func TestDay3Part1(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	want := "357"
	assertPart1(day3{}, input, want, t)
}

func TestDay3Part2(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	want := "3121910778619"
	assertPart2(day3{}, input, want, t)
}
