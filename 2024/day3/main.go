package main

import (
	"fmt"

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

func (d *Day3) Part1(input string) int {
	return -1
}

func (d *Day3) Part2(input string) int {
	return -1
}
