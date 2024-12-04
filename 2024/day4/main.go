package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

const XMAS = "XMAS"
const MAS = "MAS"

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
	diags := d.getDiags(runes, len(XMAS))
	for _, diag := range diags {
		xmas.diagonal += d.solve1(diag)
	}

	// Diagonal -- Backwards
	backwards := d.deepCopy(runes)
	for _, line := range backwards {
		slices.Reverse(line)
	}
	diags = d.getDiags(backwards, len(XMAS))
	for _, diag := range diags {
		xmas.diagonalBackwards += d.solve1(diag)
	}

	// Diagonal -- Upside-down
	upsidedown := d.deepCopy(runes)
	slices.Reverse(upsidedown)

	diags = d.getDiags(upsidedown, len(XMAS))
	for _, diag := range diags {
		xmas.diagonalUpsidedown += d.solve1(diag)
	}

	// Diagonal -- Upside-down & backwards
	for _, line := range upsidedown {
		slices.Reverse(line)
	}
	diags = d.getDiags(upsidedown, len(XMAS))
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

	indexes := map[string]int{}

	rows := len(runes)
	cols := len(runes[0])

	// Diagonal
	locs := d.getDiagsLoc(runes, len(MAS))
	for _, loc := range locs {
		for _, c := range d.solve2(loc) {
			if debug {
				fmt.Printf("%v\n", c)
			}
			d.increment(indexes, c.key())
		}
	}

	// Diagonal -- Backwards
	backwards := d.deepCopy(runes)
	for _, line := range backwards {
		slices.Reverse(line)
	}
	locs = d.getDiagsLoc(backwards, len(MAS))
	if debug {
		util.LogSuccessf("-- backwards\n")
	}
	for _, loc := range locs {
		for _, c := range d.solve2(loc) {
			// Flip x index because we flipped the whole row
			c.x = (cols + 1) - c.x

			if debug {
				fmt.Printf("%v\n", c)
			}
			d.increment(indexes, c.key())
		}
	}

	// Diagonal -- Upside-down
	upsidedown := d.deepCopy(runes)
	slices.Reverse(upsidedown)
	locs = d.getDiagsLoc(upsidedown, len(MAS))
	if debug {
		util.LogSuccessf("-- upside-down\n")
	}
	for _, loc := range locs {
		for _, c := range d.solve2(loc) {
			c.y = (rows + 1) - c.y

			if debug {
				fmt.Printf("%v\n", c)
			}
			d.increment(indexes, c.key())
		}
	}

	// Diagonal -- Upside-down & backwards
	for _, line := range upsidedown {
		slices.Reverse(line)
	}
	locs = d.getDiagsLoc(upsidedown, len(MAS))
	if debug {
		util.LogSuccessf("-- backwards and upside-down\n")
	}
	for _, loc := range locs {
		for _, c := range d.solve2(loc) {
			c.x = (cols + 1) - c.x
			c.y = (rows + 1) - c.y

			if debug {
				fmt.Printf("%v\n", c)
			}
			d.increment(indexes, c.key())
		}
	}

	result := 0
	for _, val := range indexes {
		if val == 2 {
			result++
		}
	}
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

// Return a slice of indexes of A
func (d *Day4) solve2(loc []char) []char {
	result := []char{}
	s := newStateMachine([]rune("MAS"))
	for _, char := range loc {
		if s.Next(char.rune) {
			result = append(result, char)
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

func (d *Day4) increment(kvs map[string]int, key string) {
	if _, ok := kvs[key]; !ok {
		kvs[key] = 0
	}
	kvs[key]++
}

func (d *Day4) getCol(i int, lines [][]rune) []rune {
	col := []rune{}
	for _, line := range lines {
		col = append(col, line[i])
	}
	return col
}

func (d *Day4) getDiags(lines [][]rune, minSize int) [][]rune {
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
		if len(diag) >= minSize {
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

type char struct {
	rune rune
	// x,y are 1-based indexing
	x int
	y int
}

func (c char) key() string {
	return fmt.Sprintf("%v.%v", c.x, c.y)
}

func (c char) String() string {
	return c.key() + ": " + string(c.rune)
}

func (d *Day4) getDiagsLoc(lines [][]rune, minSize int) [][]char {
	rows := len(lines)
	cols := len(lines[0])

	startY := rows - 1
	startX := 0

	chars := [][]char{}
	for {
		diag := []char{}
		x, y := startX, startY
		for {
			diag = append(diag, char{rune: rune(lines[y][x]), x: x, y: y})
			x++
			y++
			if x == cols || y == rows {
				break
			}
		}
		if len(diag) >= minSize {
			chars = append(chars, diag)
		}

		if startY > 0 {
			startY--
		} else if startX < cols-1 {
			startX++
		} else {
			break
		}
	}

	return chars
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

func printChars(chars []char) {
	runes := []rune{}
	for _, char := range chars {
		runes = append(runes, char.rune)
	}
	fmt.Printf("%v\n", string(runes))
}
