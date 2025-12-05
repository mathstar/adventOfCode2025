package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type day5 struct{}

func init() {
	registerDay(5, day5{})
}

func (_ day5) part1(input string) string {
	freshnessDb, ingredients := parseInput(input)
	fresh := countFresh(freshnessDb, ingredients)
	return strconv.Itoa(fresh)
}

func (_ day5) part2(input string) string {
	return ""
}

type freshnessEntry struct {
	start int
	fresh bool
}

func freshnessEntryCmp(a, b freshnessEntry) int {
	return a.start - b.start
}

func parseInput(input string) ([]freshnessEntry, []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var freshnessDb []freshnessEntry
	var ingredients []int
	var ingredientsSection bool

	for _, line := range lines {
		if "" == strings.TrimSpace(line) {
			ingredientsSection = true
			continue
		}

		if ingredientsSection {
			ingredient, err := strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				panic(err)
			}
			ingredients = append(ingredients, ingredient)
		} else {
			values := strings.Split(strings.TrimSpace(line), "-")
			if len(values) != 2 {
				panic("unexpected entry in freshness section")
			}

			start, err := strconv.Atoi(strings.TrimSpace(values[0]))
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(strings.TrimSpace(values[1]))
			if err != nil {
				panic(err)
			}

			freshnessDb = append(freshnessDb, freshnessEntry{
				start: start,
				fresh: true,
			})
			freshnessDb = append(freshnessDb, freshnessEntry{
				start: end + 1,
				fresh: false,
			})
		}
	}
	slices.Sort(ingredients)
	slices.SortFunc(freshnessDb, freshnessEntryCmp)

	return freshnessDb, ingredients
}

func countFresh(freshnessDb []freshnessEntry, ingredients []int) int {
	var count int
	var inFresh bool
	var nextEntry = 0

	for _, ingredient := range ingredients {
		for ; nextEntry < len(freshnessDb) && ingredient >= freshnessDb[nextEntry].start; nextEntry++ {
			inFresh = freshnessDb[nextEntry].fresh
		}

		var start, end string
		if nextEntry > 0 {
			start = strconv.Itoa(freshnessDb[nextEntry-1].start)
		} else {
			start = "0"
		}
		if nextEntry < len(freshnessDb) {
			end = strconv.Itoa(freshnessDb[nextEntry].start)
		} else {
			end = "∞"
		}
		fmt.Printf("%v: %v - %d - %v\n", inFresh, start, ingredient, end)

		if inFresh {
			count++
		}
	}

	return count
}
