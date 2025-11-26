package main

import "testing"

func TestPart1(t *testing.T) {
	input := ""
	want := "part1"
	got := day1{}.part1(input)
	if got != want {
		t.Errorf("part1() = %q, want %q", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := ""
	want := "part2"
	got := day1{}.part2(input)
	if got != want {
		t.Errorf("part2() = %q, want %q", got, want)
	}
}
