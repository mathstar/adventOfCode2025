package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	getInput(1)
	fmt.Println(days[1].part1(getInput(1)))
	fmt.Println(days[1].part2(getInput(1)))
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
