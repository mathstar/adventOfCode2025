package main

import (
	"fmt"
	"strconv"
	"strings"
)

type day6 struct{}

func init() {
	registerDay(6, day6{})
}

func (d day6) part1(input string) string {
	operands, ops := d.parseInput(input)

	var total int64
	for j := 0; j < len(operands[0]); j++ {
		o := ops[j]
		var sub int64
		if o == '*' {
			sub = 1
		}
		for i := 0; i < len(operands); i++ {
			switch o {
			case '+':
				sub += int64(operands[i][j])
			case '*':
				sub *= int64(operands[i][j])
			}
		}
		total += sub
	}

	return fmt.Sprintf("%d", total)
}

func (d day6) part2(input string) string {
	operands, ops := d.parseInput2(input)

	var total int64
	for i := 0; i < len(operands); i++ {
		o := ops[i]
		var sub int64
		if o == '*' {
			sub = 1
		}
		for _, n := range operands[i] {
			switch o {
			case '+':
				sub += int64(n)
			case '*':
				sub *= int64(n)
			}
		}
		total += sub
	}
	return fmt.Sprintf("%d", total)
}

type operand int
type op rune

func (day6) parseInput(input string) ([][]operand, []op) {
	var operands [][]operand
	var ops []op

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		split := strings.Fields(line)

		if len(split) == 0 {
			continue
		}

		_, nonNum := strconv.Atoi(split[0])
		if nonNum == nil {
			operands = append(operands, []operand{})
			for _, o := range split {
				o, err := strconv.Atoi(o)
				if err != nil {
					panic(err)
				}
				operands[len(operands)-1] = append(operands[len(operands)-1], operand(o))
			}
		} else {
			for _, o := range split {
				if o != "*" && o != "+" {
					panic(fmt.Sprintf("unexpected operator: %v", o))
				}
				ops = append(ops, op(o[0]))
			}
		}
	}
	return operands, ops
}

func (day6) parseInput2(input string) ([][]operand, []op) {
	lines := strings.Split(input, "\n")
	var chars [][]rune
	for _, line := range lines {
		if line == "" {
			continue
		}
		chars = append(chars, []rune(line))
	}
	var operands [][]operand
	operands = append(operands, []operand{})
	var ops []op
	for j := len(chars[0]) - 1; j >= 0; j-- {
		var acc strings.Builder
		for i := 0; i < len(chars); i++ {
			switch chars[i][j] {
			case ' ':
				continue
			case '*', '+':
				ops = append(ops, op(chars[i][j]))
			default:
				acc.WriteRune(chars[i][j])
			}
		}
		if acc.Len() > 0 {
			o, err := strconv.Atoi(acc.String())
			if err != nil {
				panic(err)
			}
			operands[len(operands)-1] = append(operands[len(operands)-1], operand(o))
			if len(operands) == len(ops) {
				operands = append(operands, []operand{})
			}
		}
	}
	return operands[:len(ops)], ops
}
