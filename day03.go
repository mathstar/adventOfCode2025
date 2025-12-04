package main

import (
	"strconv"
	"strings"
)

type day3 struct{}

func init() {
	registerDay(3, day3{})
}

func (_ day3) part1(input string) string {
	var totalJoltage int
	for _, bank := range parseBatteries(input) {
		totalJoltage += maximizeJoltage(bank)
	}
	return strconv.Itoa(totalJoltage)
}

func (_ day3) part2(input string) string {
	banks := parseBatteries(input)
	channels := make([]chan int, len(banks))
	for i := range channels {
		channels[i] = make(chan int)
	}
	for i, bank := range parseBatteries(input) {
		go maximizeLargeJoltage(bank, channels[i])
	}

	var totalJoltage int
	for _, ch := range channels {
		totalJoltage += <-ch
	}
	return strconv.Itoa(totalJoltage)
}

func parseBatteries(input string) [][]int {
	var batteries [][]int
	for bank := range strings.Lines(input) {
		var batteryValues []int
		for _, battery := range strings.TrimSpace(bank) {
			batteryValues = append(batteryValues, int(battery-'0'))
		}
		batteries = append(batteries, batteryValues)
	}
	return batteries
}

func maximizeJoltage(bank []int) int {
	maxJoltage, maxToRight := 0, bank[len(bank)-1]
	for i := len(bank) - 2; i >= 0; i-- {
		left := bank[i]
		maxJoltage = max(maxJoltage, left*10+maxToRight)
		maxToRight = max(maxToRight, left)
	}
	return maxJoltage
}

//func maximizeLargeJoltage(bank []int) int {
//	enabled := make([]bool, len(bank))
//	placements := make([][]int, 10)
//	for i := 0; i < 12; i++ {
//		enabled[i] = true
//		placements[bank[i]] = append(placements[bank[i]], i)
//	}
//
//outer:
//	for i := 12; i < len(bank); i++ {
//		for v := 0; v <= bank[i]; v++ {
//			if len(placements[v]) > 0 {
//				j := placements[v][0]
//
//				enabled[i], enabled[j] = true, false
//				placements[v] = placements[v][1:]
//				placements[bank[i]] = append(placements[bank[i]], i)
//				continue outer
//			}
//		}
//	}
//
//	var total int
//	for i := 0; i < len(bank); i++ {
//		if enabled[i] {
//			total *= 10
//			total += bank[i]
//		}
//	}
//	return total
//}

//func maximizeLargeJoltage(bank []int) int {
//	digits := []int{bank[0], bank[1], bank[2], bank[3], bank[4], bank[5], bank[6], bank[7], bank[8], bank[9], bank[10], bank[11]}
//
//outer:
//	for i := 12; i < len(bank); i++ {
//		//fmt.Printf("%v\n", digits)
//		nextDigit := bank[i]
//		for j, digit := range digits {
//			if j < len(digits)-1 && digit > digits[j+1] {
//				continue
//			}
//			if nextDigit >= digit {
//				//fmt.Printf("b[%v]=%v -> b[%v]=%v\n", j, digit, i, nextDigit)
//				digits = append(digits[:j], digits[j+1:]...)
//				digits = append(digits, nextDigit)
//				continue outer
//			}
//		}
//	}
//	//fmt.Printf("%v\n", digits)
//	//fmt.Println()
//
//	var total int
//	for _, n := range digits {
//		total *= 10
//		total += n
//	}
//	return total
//}

func maximizeLargeJoltage(bank []int, ch chan int) {
	memo := make(map[memoKey]int)
	ch <- build(&memo, &bank, 0, 0, 0)
}

type memoKey struct {
	value      int
	digitCount int
	i          int
}

func build(memo *map[memoKey]int, bank *[]int, value, digitCount, i int) int {
	key := memoKey{value, digitCount, i}
	if (*memo)[key] > 0 {
		return (*memo)[key]
	}

	if i >= len(*bank) {
		return value
	}

	if digitCount < 12 {
		return build(memo, bank, value*10+(*bank)[i], digitCount+1, i+1)
	}

	mask := 10
	nextValue := value
	for j := 0; j < 12; j++ {
		valueWithDigitSwap := (value%(mask/10)*10 + value/mask*mask) + (*bank)[i]
		nextValue = max(nextValue, valueWithDigitSwap)
		mask *= 10
	}
	m := build(memo, bank, nextValue, digitCount, i+1)
	(*memo)[key] = m
	return m
}
