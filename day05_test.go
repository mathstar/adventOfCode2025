package main

import "testing"

func TestDay5Part1(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	want := "3"
	assertPart1(day5{}, input, want, t)
}

func TestDay5Part2(t *testing.T) {
	input := ``
	want := ""
	assertPart2(day5{}, input, want, t)
}
