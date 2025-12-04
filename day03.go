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

func maximizeLargeJoltage(bank []int, ch chan int) {
	// start with first 12 digits
	var value int
	for i := 0; i < 12; i++ {
		value *= 10
		value += bank[i]
	}

	for i := 12; i < len(bank); i++ {
		// for each additional digit, find maximum value of either eliminating an existing digit in favor of the new digit
		// or keeping current value
		mask := 10
		startingValue := value
		for j := 0; j < 12; j++ {
			valueWithDigitSwap := (startingValue%(mask/10)*10 + startingValue/mask*mask) + bank[i]
			value = max(value, valueWithDigitSwap)
			mask *= 10
		}
	}

	ch <- value
}
