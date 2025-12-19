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
	m := d.findMaxConstrainedRectangle(d.parseInput(input))
	return strconv.Itoa(m)
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
			xa, xb := min(pairs[i][0], pairs[j][0]), max(pairs[i][0], pairs[j][0])
			ya, yb := min(pairs[i][1], pairs[j][1]), max(pairs[i][1], pairs[j][1])
			size := xb - xa + 1
			size *= yb - ya + 1
			m = max(m, size)
		}
	}
	return m
}

func (d day9) findMaxConstrainedRectangle(pairs [][]int) int {
	var m int
	for i := range pairs {
		for j := i + 1; j < len(pairs); j++ {
			if !d.isUsableRectangle(pairs, pairs[i], pairs[j]) {
				continue
			}
			xa, xb := min(pairs[i][0], pairs[j][0]), max(pairs[i][0], pairs[j][0])
			ya, yb := min(pairs[i][1], pairs[j][1]), max(pairs[i][1], pairs[j][1])
			size := xb - xa + 1
			size *= yb - ya + 1
			m = max(m, size)
		}
	}
	return m
}

func (d day9) isUsableRectangle(corners [][]int, a, b []int) bool {
	return !d.rectangleContainsCorner(corners, a, b) && d.rectangleInteriorIsInsidePolygon(corners, a, b)
}

func (day9) rectangleContainsCorner(corners [][]int, a, b []int) bool {
	xa, xb := min(a[0], b[0]), max(a[0], b[0])
	ya, yb := min(a[1], b[1]), max(a[1], b[1])
	for _, c := range corners {
		if (c[0] == a[0] && c[1] == a[1]) || (c[0] == b[0] && c[1] == b[1]) {
			continue
		}

		if c[0] > xa && c[0] < xb && c[1] > ya && c[1] < yb {
			// point is within the rectangle
			return true
		}

		// TODO logic not quite right, corner on edge should only disqualify if its edges enter interior of polygon
		if c[0] == xa || c[0] == xb {
			if c[1] > ya && c[1] < yb {
				// point is along edge of rectangle
				return true
			}
		}

		if c[1] == ya || c[1] == yb {
			if c[0] > xa && c[0] < xb {
				// point is along edge of rectangle
				return true
			}
		}
	}
	return false
}

func (day9) rectangleInteriorIsInsidePolygon(corners [][]int, a, b []int) bool {
	xa, xb := min(a[0], b[0]), max(a[0], b[0])
	ya, yb := min(a[1], b[1]), max(a[1], b[1])

	x, y := (xb+xa)/2, (yb+ya)/2

	var crossings int
	for i := range corners {
		if i+1 < len(corners) {
			j := i + 1
			// find horizontal edges along ray from point to y=0
			if corners[i][1] == corners[j][1] {
				// is horizontal
				if corners[i][1] <= y {
					// is along vertical path of ray
					if (corners[i][0] < x && corners[j][0] > x) || (corners[i][0] > x && corners[j][0] < x) {
						// intersects ray
						crossings++
					}
				}
			}
		}
	}

	// odd number of crossings indicates interior of polygon
	return crossings%2 == 1
}
