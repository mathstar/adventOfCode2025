package main

import "testing"

func TestDay9Part1(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	want := "50"
	assertPart1(day9{}, input, want, t)
}

func TestDay9Part2(t *testing.T) {
	input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	want := "24"
	assertPart2(day9{}, input, want, t)
}
