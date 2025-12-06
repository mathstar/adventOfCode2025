package main

import "testing"

func TestDay6Part1(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`
	want := "4277556"
	assertPart1(day6{}, input, want, t)
}

func TestDay6Part2(t *testing.T) {
	input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`
	want := "3263827"
	assertPart2(day6{}, input, want, t)
}
