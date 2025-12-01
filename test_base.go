package main

import "testing"

func assertPart1(d day, input, want string, t *testing.T) {
	got := d.part1(input)
	if got != want {
		t.Errorf("part1() = %q, want %q", got, want)
	}
}

func assertPart2(d day, input, want string, t *testing.T) {
	got := d.part2(input)
	if got != want {
		t.Errorf("part2() = %q, want %q", got, want)
	}
}
