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
	results := make([]chan int64, len(ranges))
	for i, r := range ranges {
		results[i] = make(chan int64)
		go determineInvalidCount(r, results[i])
	}
	var sum int64
	for _, r := range results {
		sum += <-r
	}
	return strconv.FormatInt(sum, 10)
}

func determineInvalidCount(r string, ch chan int64) {
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

func invalidSum(startInc, endInc int) int64 {
	var sum int64
	for i := startInc; i <= endInc; i++ {
		s := strconv.Itoa(i)
		if len(s)%2 == 1 {
			continue
		}
		if s[0:len(s)/2] == s[len(s)/2:] {
			sum += int64(i)
		}
	}
	return sum
}

func (_ day2) part2(input string) string {
	return ""
}
