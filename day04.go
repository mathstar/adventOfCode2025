package main

import (
	"strconv"
	"strings"
)

type day4 struct{}

func init() {
	registerDay(4, day4{})
}

func (day4) part1(input string) string {
	grid := parseGrid(input)

	var movable int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] && neighborCount(&grid, i, j) < 4 {
				movable++
			}
		}
	}

	return strconv.Itoa(movable)
}

func (day4) part2(input string) string {
	grid := parseGrid(input)

	var movable int
	for removed := true; removed; {
		removed = false
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] && neighborCount(&grid, i, j) < 4 {
					removed = true
					movable++
					grid[i][j] = false
				}
			}
		}
	}

	return strconv.Itoa(movable)
}

func parseGrid(input string) [][]bool {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]bool, len(lines))
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		grid[i] = make([]bool, len(trimmed))
		for j, char := range trimmed {
			grid[i][j] = char == '@'
		}
	}
	return grid
}

func neighborCount(grid *[][]bool, i, j int) int {
	var count int
	if i > 0 && j > 0 && (*grid)[i-1][j-1] {
		count++
	}
	if i > 0 && (*grid)[i-1][j] {
		count++
	}
	if i > 0 && j < len((*grid)[i-1])-1 && (*grid)[i-1][j+1] {
		count++
	}

	if j > 0 && (*grid)[i][j-1] {
		count++
	}
	if j < len((*grid)[i])-1 && (*grid)[i][j+1] {
		count++
	}

	if i < len(*grid)-1 && j > 0 && (*grid)[i+1][j-1] {
		count++
	}
	if i < len(*grid)-1 && (*grid)[i+1][j] {
		count++
	}
	if i < len(*grid)-1 && j < len((*grid)[i+1])-1 && (*grid)[i+1][j+1] {
		count++
	}

	return count
}
