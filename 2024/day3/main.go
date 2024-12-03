package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
)

var debug bool = false

func main() {
	day := &Day3{}
	input := day.Parse()

	debug = len(os.Args) == 3 && os.Args[2] == "--debug"

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
	MUL_PATTERN  = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	DO_PATTERN   = regexp.MustCompile(`do\(\)`)
	DONT_PATTERN = regexp.MustCompile(`don't\(\)`)
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
	result := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		result += d.findMulWithDos(line)
	}
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

/* TODO(joeriddles): Use a state machine */
func (d *Day3) findMulWithDos(input string) int {
	result := 0

	dos := d.getIndexes(DO_PATTERN.FindAllStringSubmatchIndex(input, -1))
	donts := d.getIndexes(DONT_PATTERN.FindAllStringSubmatchIndex(input, -1))
	muls := MUL_PATTERN.FindAllStringSubmatchIndex(input, -1)

	if debug {
		d.printMatches(input, dos, donts, muls)
	}

	for _, match := range muls {
		enabled := d.isEnabled(match[0], &dos, &donts)
		if enabled {
			lstr := input[match[2]:match[3]]
			rstr := input[match[4]:match[5]]
			l, _ := strconv.Atoi(lstr)
			r, _ := strconv.Atoi(rstr)
			result += l * r
		}
	}

	return result
}

// Check if the most recent command is to enable or disable multiplication.
func (d *Day3) isEnabled(i int, dos *[]int, donts *[]int) bool {
	lastDo := -1
	lastDont := -1

	for _, j := range *dos {
		if j >= i {
			break
		}
		lastDo = j
	}

	for _, j := range *donts {
		if j >= i {
			break
		}
		lastDont = j
	}

	return lastDo == -1 && lastDont == -1 || lastDo > lastDont
}

func (d *Day3) getIndexes(matches [][]int) []int {
	indexes := []int{}
	for _, match := range matches {
		indexes = append(indexes, match[0])
	}
	return indexes
}

const (
	Reset string = "\033[0m"
	Red   string = "\033[31m"
	Green string = "\033[32m"
)

func (d *Day3) printMatches(input string, dos []int, donts []int, muls [][]int) {
	mulIndexes := d.getIndexes(muls)

	result := 0
	enabled := true
	for i := range input {
		if slices.Contains(dos, i) {
			enabled = true
			fmt.Printf(Reset+"\n%4v: %v -- %v\n", i, input[i:i+4], result)
		}
		if slices.Contains(donts, i) {
			enabled = false
			fmt.Printf(Reset+"\n%4v: %v -- %v\n", i, input[i:i+7], result)
		}
		if slices.Contains(mulIndexes, i) {
			for _, mul := range muls {
				color := Green
				if !enabled {
					color = Red
				}

				if mul[0] == i {
					if enabled {
						lstr := input[mul[2]:mul[3]]
						rstr := input[mul[4]:mul[5]]
						l, _ := strconv.Atoi(lstr)
						r, _ := strconv.Atoi(rstr)
						result += l * r
					}

					fmt.Printf(color+"%v: %v ", i, input[mul[0]:mul[1]])
					break
				}
			}
		}
	}
	fmt.Println(Reset + "\n-------")
}
