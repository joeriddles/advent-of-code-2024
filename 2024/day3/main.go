package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
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

var MUL_PATTERN = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func (d *Day3) Part1(input string) int {
	result := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		result += d.findMul(line)
	}
	return result
}

func (d *Day3) Part2(input string) int {
	result := 0

	return result
}

func (d *Day3) findMul(input string) int {
	result := 0
	matches := MUL_PATTERN.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		lstr := match[1]
		rstr := match[2]
		l, _ := strconv.Atoi(lstr)
		r, _ := strconv.Atoi(rstr)
		result += l * r
	}
	return result
}
