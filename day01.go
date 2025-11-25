package main

type day1 struct{}

func init() {
	registerDay(1, day1{})
}

func (d day1) part1(input string) string {
	return "part1"
}

func (d day1) part2(input string) string {
	return "part2"
}
