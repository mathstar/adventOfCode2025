package main

import (
	"strconv"
	"strings"
)

type day2 struct{}

func init() {
	registerDay(2, day2{})
}

func (_ day2) part1(input string) string {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	results := make([]chan int, len(ranges))
	for i, r := range ranges {
		results[i] = make(chan int)
		go determineInvalidCount(r, results[i])
	}
	var sum int
	for _, r := range results {
		sum += <-r
	}
	return strconv.Itoa(sum)
}

func determineInvalidCount(r string, ch chan int) {
	split := strings.Split(r, "-")
	startInc, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	endInc, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	ch <- invalidSum(startInc, endInc)
}

func invalidSum(startInc, endInc int) int {
	var sum int
	for i := startInc; i <= endInc; i++ {
		s := strconv.Itoa(i)
		if len(s)%2 == 1 {
			continue
		}
		if s[0:len(s)/2] == s[len(s)/2:] {
			sum += i
		}
	}
	return sum
}

func (_ day2) part2(input string) string {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	results := make([]chan int, len(ranges))
	for i, r := range ranges {
		results[i] = make(chan int)
		go determineInvalidCount2(r, results[i])
	}
	var sum int
	for _, r := range results {
		sum += <-r
	}
	return strconv.Itoa(sum)
}

func determineInvalidCount2(r string, ch chan int) {
	split := strings.Split(r, "-")
	startInc, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	endInc, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	ch <- invalidSum2(startInc, endInc)
}

func invalidSum2(startInc, endInc int) int {
	var sum int
numberLoop:
	for i := startInc; i <= endInc; i++ {
		s := strconv.Itoa(i)
	splitLengthLoop:
		for splitLength := 1; splitLength <= len(s)/2; splitLength++ {
			if len(s)%splitLength == 0 {
				var splits []string
				for j := 0; j < len(s); j += splitLength {
					splits = append(splits, s[j:j+splitLength])
				}
				for j := 1; j < len(splits); j++ {
					if splits[0] != splits[j] {
						continue splitLengthLoop
					}
				}
				sum += i
				continue numberLoop
			}
		}
	}
	return sum
}
