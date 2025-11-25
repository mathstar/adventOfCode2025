package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var selectedDay int

func init() {
	flag.IntVar(&selectedDay, "day", 0, "Day to run")
	flag.IntVar(&selectedDay, "d", 0, "Day to run")
}

func main() {
	flag.Parse()
	if selectedDay != 0 {
		runDay(selectedDay)
	} else {
		for i := range days {
			runDay(i)
		}
	}
}

func runDay(day int) {
	dayImpl, ok := days[day]
	if !ok {
		panic(fmt.Sprintf("Day %02d not registered", day))
	}

	fmt.Printf("Running day %02d\n", day)
	input := getInput(day)
	start := time.Now()
	fmt.Println(dayImpl.part1(input))
	part1Time := time.Since(start)
	start = time.Now()
	fmt.Println(dayImpl.part2(input))
	part2Time := time.Since(start)
	fmt.Printf("Timing: part1 %s, part2 %s, total %s\n", part1Time, part2Time, part1Time+part2Time)
}

type day interface {
	part1(input string) string
	part2(input string) string
}

var days = make(map[int]day)

func registerDay(day int, d day) {
	days[day] = d
}

func getInput(day int) string {
	// Use cached input if available
	cached, err := os.ReadFile(fmt.Sprintf("cache/day%02d.txt", day))
	if err == nil {
		return string(cached)
	}

	// Ensure cache directory exists
	err = os.Mkdir("cache", 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	// Fetch input from Advent of Code
	userAgent := "https://github.com/mathstar/adventOfCode2025 by mstaricka@gmail.com"
	cookie := "session=" + os.Getenv("SESSION")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day), nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", cookie)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("cache/day%02d.txt", day), body, 0644)
	if err != nil {
		panic(err)
	}

	return string(body)
}
