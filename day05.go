package main

import (
	"slices"
	"strconv"
	"strings"
)

type day5 struct{}

func init() {
	registerDay(5, day5{})
}

func (day5) part1(input string) string {
	freshnessDb, ingredients := parseInput(input)
	fresh := countFresh(freshnessDb, ingredients)
	return strconv.Itoa(fresh)
}

func (day5) part2(input string) string {
	freshnessDb, _ := parseInput(input)
	fresh := countAllFresh(freshnessDb)
	return strconv.Itoa(fresh)
}

type freshnessEntry struct {
	start int
	end   int
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

			freshnessDb = append(freshnessDb, freshnessEntry{start, end})
		}
	}
	slices.Sort(ingredients)
	slices.SortFunc(freshnessDb, freshnessEntryCmp)

	return freshnessDb, ingredients
}

func countFresh(freshnessDb []freshnessEntry, ingredients []int) int {
	var count int
	var searchStart int

	for _, ingredient := range ingredients {
		for ; searchStart < len(freshnessDb) && ingredient > freshnessDb[searchStart].end; searchStart++ {
		}
		if searchStart >= len(freshnessDb) {
			break
		}

		if ingredient >= freshnessDb[searchStart].start && ingredient <= freshnessDb[searchStart].end {
			count++
		}
	}

	return count
}

func countAllFresh(freshnessDb []freshnessEntry) int {
	var count, p int
	for _, entry := range freshnessDb {
		if p > entry.end {
			continue
		}
		start := max(entry.start, p)
		end := entry.end

		count += end - start + 1
		p = end + 1
	}
	return count
}
