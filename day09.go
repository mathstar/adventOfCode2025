package main

import (
	"strconv"
	"strings"
)

type day9 struct{}

func init() {
	registerDay(9, day9{})
}

func (d day9) part1(input string) string {
	m := d.findMaxRectangle(d.parseInput(input))
	return strconv.Itoa(m)
}

func (d day9) part2(input string) string {
	return ""
}

func (day9) parseInput(input string) [][]int {
	var pairs [][]int
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		split := strings.Split(line, ",")
		a, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		pairs = append(pairs, []int{a, b})
	}
	return pairs
}

func (day9) findMaxRectangle(pairs [][]int) int {
	var m int
	for i := range pairs {
		for j := i + 1; j < len(pairs); j++ {
			size := pairs[i][0] - pairs[j][0] + 1
			size *= pairs[i][1] - pairs[j][1] + 1
			if size < 0 {
				size *= -1
			}
			m = max(m, size)
		}
	}
	return m
}
