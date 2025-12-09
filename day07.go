package main

import (
	"fmt"
	"strings"
)

type day7 struct{}

func init() {
	registerDay(7, day7{})
}

func (d day7) part1(input string) string {
	var beams []int
	var splitCount int
	for _, line := range d.parseInput(input) {
		var stepSplitCount int
		beams, stepSplitCount = d.stepBeam(line, beams)
		splitCount += stepSplitCount
	}
	return fmt.Sprintf("%d", splitCount)
}

func (d day7) part2(input string) string {
	timelines := make(map[int]int)
	var init bool
	for _, line := range d.parseInput(input) {
		if !init {
			beams, _ := d.stepBeam(line, []int{})
			for _, beam := range beams {
				timelines[beam]++
			}
			init = true
		} else {
			stepTimelines := make(map[int]int)
			for beam, count := range timelines {
				steppedBeams, _ := d.stepBeam(line, []int{beam})
				for _, steppedBeam := range steppedBeams {
					stepTimelines[steppedBeam] += count
				}
			}
			timelines = stepTimelines
		}
	}

	timelineCount := 0
	for _, count := range timelines {
		timelineCount += count
	}
	return fmt.Sprintf("%d", timelineCount)
}

func (day7) parseInput(input string) [][]rune {
	var result [][]rune
	for line := range strings.Lines(input) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, []rune(line))
		result[len(result)-1] = []rune(line)
	}
	return result
}

func (day7) stepBeam(line []rune, beams []int) ([]int, int) {
	var nextBeams []int
	var splitCount int
	for i, c := range line {
		if c == 'S' {
			nextBeams = append(nextBeams, i)
		}
	}
	for _, i := range beams {
		switch line[i] {
		case '.':
			if len(nextBeams) == 0 || nextBeams[len(nextBeams)-1] != i {
				nextBeams = append(nextBeams, i)
			}
		case '^':
			splitCount++
			if len(nextBeams) == 0 || nextBeams[len(nextBeams)-1] != i-1 {
				nextBeams = append(nextBeams, i-1)
			}
			nextBeams = append(nextBeams, i+1)
		}
	}
	return nextBeams, splitCount
}
