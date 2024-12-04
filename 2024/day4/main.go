package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
)

var debug = false

func main() {
	day := &Day4{}
	input := day.Parse()

	answer1 := day.Part1(input)
	fmt.Printf("part1: %v\n", answer1)

	answer2 := day.Part2(input)
	fmt.Printf("part2: %v\n", answer2)
}

type Day4 struct {
	*day.BaseDay
}

func NewDay4() day.Day {
	return &Day4{
		BaseDay: day.NewBaseDay(),
	}
}

var WHITESPACE = regexp.MustCompile(`\s+`)

type xmas struct {
	horizontal                  int
	horizontalBackwards         int
	vertical                    int
	verticalUpsidedown          int
	diagonal                    int
	diagonalBackwards           int
	diagonalUpsidedown          int
	diagonalBackwardsUpsidedown int
}

func (x *xmas) sum() int {
	return x.horizontal +
		x.horizontalBackwards +
		x.vertical +
		x.verticalUpsidedown +
		x.diagonal +
		x.diagonalBackwards +
		x.diagonalUpsidedown +
		x.diagonalBackwardsUpsidedown
}

func (d *Day4) Part1(input string) int {
	lines := strings.Split(input, "\n")
	runes := [][]rune{}
	for _, line := range lines {
		if line == "" || WHITESPACE.MatchString(line) {
			continue
		}
		runes = append(runes, []rune(line))
	}

	xmas := &xmas{}

	// Horizontal
	for i, line := range runes {
		c := d.solve1(line)
		xmas.horizontal += c
		if debug {
			fmt.Printf("horizontal %v: %v -- %v\n", i, string(line), c)
		}

		// Backwards
		backwards := slices.Clone(line)
		slices.Reverse(backwards)
		c = d.solve1(backwards)
		xmas.horizontalBackwards += c
		if debug {
			fmt.Printf("backwards  %v: %v -- %v\n", i, string(backwards), c)
		}
	}

	// Vertical -- assume all lines are same length
	for i := range runes[0] {
		col := d.getCol(i, runes)
		xmas.vertical += d.solve1(col)

		// Backwards
		backwards := slices.Clone(col)
		slices.Reverse(backwards)
		xmas.verticalUpsidedown += d.solve1(backwards)
	}

	// Diagonal
	diags := d.getDiags(runes)
	for _, diag := range diags {
		xmas.diagonal += d.solve1(diag)
	}

	// Diagonal -- Backwards
	backwards := d.deepCopy(runes)
	for _, line := range backwards {
		slices.Reverse(line)
	}
	diags = d.getDiags(backwards)
	for _, diag := range diags {
		xmas.diagonalBackwards += d.solve1(diag)
	}

	// Diagonal -- Upside-down
	upsidedown := d.deepCopy(runes)
	slices.Reverse(upsidedown)

	diags = d.getDiags(upsidedown)
	for _, diag := range diags {
		xmas.diagonalUpsidedown += d.solve1(diag)
	}

	// Diagonal -- Upside-down & backwards
	for _, line := range upsidedown {
		slices.Reverse(line)
	}
	diags = d.getDiags(upsidedown)
	for _, diag := range diags {
		xmas.diagonalBackwardsUpsidedown += d.solve1(diag)
	}

	result := xmas.sum()
	return result
}

func (d *Day4) Part2(input string) int {
	lines := strings.Split(input, "\n")
	runes := [][]rune{}
	for _, line := range lines {
		if line == "" || WHITESPACE.MatchString(line) {
			continue
		}
		runes = append(runes, []rune(line))
	}

	xmas := &xmas{}

	// Diagonal
	diags := d.getDiags(runes)
	for _, diag := range diags {
		xmas.diagonal += d.solve1(diag)
	}

	// Diagonal -- Backwards
	backwards := d.deepCopy(runes)
	for _, line := range backwards {
		slices.Reverse(line)
	}
	diags = d.getDiags(backwards)
	for _, diag := range diags {
		xmas.diagonalBackwards += d.solve1(diag)
	}

	result := xmas.sum()
	return result
}

func (d *Day4) solve1(str []rune) int {
	result := 0
	s := newStateMachine([]rune("XMAS"))
	for _, char := range str {
		if s.Next(char) {
			result++
		}
	}
	return result
}

func (d *Day4) deepCopy(runes [][]rune) [][]rune {
	copy := [][]rune{}
	for _, line := range runes {
		copy = append(copy, slices.Clone(line))
	}
	return copy
}

func (d *Day4) getCol(i int, lines [][]rune) []rune {
	col := []rune{}
	for _, line := range lines {
		col = append(col, line[i])
	}
	return col
}

func (d *Day4) getDiags(lines [][]rune) [][]rune {
	rows := len(lines)
	cols := len(lines[0])

	startY := rows - 1
	startX := 0

	diags := [][]rune{}
	for {
		diag := []rune{}
		x, y := startX, startY
		for {
			diag = append(diag, rune(lines[y][x]))
			x++
			y++
			if x == cols || y == rows {
				break
			}
		}
		if len(diag) >= 4 {
			diags = append(diags, diag)
		}

		if startY > 0 {
			startY--
		} else if startX < cols-1 {
			startX++
		} else {
			break
		}
	}

	return diags
}

func newStateMachine(word []rune) *stateMachine {
	return &stateMachine{
		word:  word,
		cur:   rune(0),
		index: -1,
	}
}

type stateMachine struct {
	word  []rune
	cur   rune
	index int
}

func (s *stateMachine) Next(next rune) bool {
	if next == s.word[0] { // This assumes no repeating characters...
		s.cur = next
		s.index = 0
	} else if s.index >= 0 && s.cur == s.word[s.index] && next == s.word[s.index+1] {
		s.cur = next
		s.index++
		if s.index == len(s.word)-1 {
			s.cur = rune(0)
			s.index = -1
			return true
		}
	} else {
		s.cur = rune(0)
		s.index = -1
	}
	return false
}
