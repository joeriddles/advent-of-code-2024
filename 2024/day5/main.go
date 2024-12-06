package main

import (
	"fmt"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
)

func main() {
	day := &Day5{}
	input := day.Parse()

	answer1 := day.Part1(input)
	fmt.Printf("part1: %v\n", answer1)

	answer2 := day.Part2(input)
	fmt.Printf("part2: %v\n", answer2)
}

type Day5 struct {
	*day.BaseDay
}

func NewDay5() day.Day {
	return &Day5{
		BaseDay: day.NewBaseDay(),
	}
}

func (d *Day5) Part1(input string) int {
	return -1
}

func (d *Day5) Part2(input string) int {
	return -1
}
