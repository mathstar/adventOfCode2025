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
		go determineInvalidCount(r, checkValiditySimple, results[i])
	}
	var sum int
	for _, r := range results {
		sum += <-r
	}
	return strconv.Itoa(sum)
}

func (_ day2) part2(input string) string {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	results := make([]chan int, len(ranges))
	for i, r := range ranges {
		results[i] = make(chan int)
		go determineInvalidCount(r, checkValidityComplex, results[i])
	}
	var sum int
	for _, r := range results {
		sum += <-r
	}
	return strconv.Itoa(sum)
}

func determineInvalidCount(r string, validityFunc func(int) bool, ch chan int) {
	split := strings.Split(r, "-")
	startInc, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	endInc, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	ch <- invalidSum(startInc, endInc, validityFunc)
}

func invalidSum(startInc, endInc int, validityFunc func(int) bool) int {
	var sum int
	for i := startInc; i <= endInc; i++ {
		if validityFunc(i) {
			sum += i
		}
	}
	return sum
}

func checkValiditySimple(i int) bool {
	s := strconv.Itoa(i)
	if len(s)%2 == 1 {
		return false
	}
	if s[0:len(s)/2] == s[len(s)/2:] {
		return true
	}
	return false
}

func checkValidityComplex(i int) bool {
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
			return true
		}
	}
	return false
}
