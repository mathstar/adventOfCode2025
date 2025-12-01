package main

import (
	"fmt"
	"strconv"
	"strings"
)

type day1 struct{}

func init() {
	registerDay(1, day1{})
}

func (d day1) part1(input string) string {
	zeroCount := 0
	position := 50
	for l := range strings.Lines(input) {
		distance, err := strconv.Atoi(strings.TrimSpace(l[1:]))
		if err != nil {
			panic(err)
		}

		switch l[0] {
		case 'L':
			position -= distance
		case 'R':
			position += distance
		}
		position = (position + 100) % 100
		if position == 0 {
			zeroCount++
		}
	}
	return fmt.Sprintf("%d", zeroCount)
}

func (d day1) part2(input string) string {
	clickCount := 0
	position := 50
	for l := range strings.Lines(input) {
		distance, err := strconv.Atoi(strings.TrimSpace(l[1:]))
		if err != nil {
			panic(err)
		}

		multiplier := 1
		distToZero := 0
		switch l[0] {
		case 'L':
			multiplier = -1
			distToZero = position
		case 'R':
			multiplier = 1
			distToZero = 100 - position
		}
		if distToZero == 0 {
			// require full rotation to increment click count if already at zero
			distToZero = 100
		}

		for distance >= distToZero {
			clickCount++
			distance -= distToZero
			position = 0
			distToZero = 100
		}

		position = (position + distance*multiplier + 100) % 100
	}
	return fmt.Sprintf("%d", clickCount)
}
