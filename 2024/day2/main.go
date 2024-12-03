package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg"
)

func main() {
	day := &Day2{}
	input := day.Parse()

	answer1 := day.Part1(input)
	fmt.Printf("part1: %v\n", answer1)

	answer2 := day.Part2(input)
	fmt.Printf("part2: %v\n", answer2)
}

type Day2 struct {
	*pkg.BaseDay
}

func NewDay1() pkg.Day {
	return &Day2{
		BaseDay: pkg.NewBaseDay(),
	}
}

const INCREASING = true

func (d *Day2) Part1(input string) int {
	reports := d.parse(input)
	safe := 0
	for _, report := range reports {
		if d.isSafe(report) {
			safe++
		}
	}
	return safe
}

func (d *Day2) Part2(input string) int {
	reports := d.parse(input)
	safe := 0
	for _, report := range reports {
		if d.isAlmostSafe(report) {
			safe++
		}
	}
	return safe
}

func (d *Day2) isSafe(report *[]int) bool {
	var increasing *bool = nil
	prev := (*report)[0]
	for _, level := range (*report)[1:] {
		diff := prev - level
		prev = level

		inc := diff < 0
		if increasing == nil {
			increasing = &inc
		}
		if inc != *increasing {
			return false
		}

		a := abs(diff)
		if a < 1 || a > 3 {
			return false
		}
	}
	return true
}

// If the report would be safe if one value was removed
func (d *Day2) isAlmostSafe(report *[]int) bool {
	isSafe := d.isSafe(report)
	if isSafe {
		return true
	}

	rep := *report
	for i := range rep {
		copy := append(slices.Clone(rep[:i]), rep[i+1:]...)
		isSafe := d.isSafe(&copy)
		if isSafe {
			return true
		}
	}
	return false

}

func (d *Day2) parse(input string) []*[]int {
	reports := []*[]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		report := []int{}
		for _, part := range parts {
			level, _ := strconv.Atoi(part)
			report = append(report, level)
		}
		reports = append(reports, &report)
	}
	return reports
}

func abs(i int) int {
	if i < 0 {
		i = -1 * i
	}
	return i
}
