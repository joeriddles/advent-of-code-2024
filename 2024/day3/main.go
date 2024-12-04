package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

func main() {
	day := &Day3{}
	input := day.Parse()

	answer1 := day.Part1(input)
	fmt.Printf("part1: %v\n", answer1)

	answer2 := day.Part2(input)
	fmt.Printf("part2: %v\n", answer2)
}

type Day3 struct {
	*day.BaseDay
}

func NewDay3() day.Day {
	return &Day3{
		BaseDay: day.NewBaseDay(),
	}
}

var (
	MUL_PATTERN = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	ALL_PATTERN = regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\))|don't\(\)`)
)

func (d *Day3) Part1(input string) int {
	result := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		result += d.findMul(line)
	}
	return result
}

func (d *Day3) Part2(input string) int {
	return d.findMulWithDos(input)
}

func (d *Day3) findMul(input string) int {
	result := 0
	matches := MUL_PATTERN.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		var l, r int
		_, err := fmt.Sscanf(match[0], "mul(%d,%d)", &l, &r)
		if err != nil {
			util.LogErr(err)
			continue
		}
		result += l * r
	}
	return result
}

func (d *Day3) findMulWithDos(input string) int {
	result := 0

	enabled := true
	for _, match := range ALL_PATTERN.FindAllString(input, -1) {
		switch match {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if !enabled {
				continue
			}

			var l, r int
			_, err := fmt.Sscanf(match, "mul(%d,%d)", &l, &r)
			if err != nil {
				util.LogErr(err)
				continue
			}

			result += l * r
		}
	}

	return result
}
