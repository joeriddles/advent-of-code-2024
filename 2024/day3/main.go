package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

const (
	Reset string = "\033[0m"
	Red   string = "\033[31m"
	Green string = "\033[32m"
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
	ALL_PATTERN  = regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\))|don't\(\)`)
)

type multiply struct {
	l int
	r int
}
type do struct{}
type dont struct{}

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
		l := d.parseInt(lstr)
		r := d.parseInt(rstr)
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

	// iter := d.makeIter(input, dos, donts, muls)
	// if debug {
	// 	d.printIter(iter)
	// }
	// enabled := true
	// for _, item := range iter {
	// 	switch i := item.(type) {
	// 	case do:
	// 		enabled = true
	// 	case dont:
	// 		enabled = false
	// 	case multiply:
	// 		if enabled {
	// 			result += (i.l * i.r)
	// 		}
	// 	}
	// }

	return result
}

const MaxInt = int(^uint(0) >> 1)

// Make a slice that contains, in order, the do's, don'ts, and multiplication's.
func (d *Day3) makeIter(input string, dos []int, donts []int, muls [][]int) []any {
	iter := []any{}
	empty := false
	for {
		doi := util.HeadOrDefault(dos, MaxInt)
		donti := util.HeadOrDefault(donts, MaxInt)
		muli := util.HeadOrDefault(muls, []int{MaxInt})[0]

		if doi < donti && doi < muli {
			iter = append(iter, do{})
			dos = dos[1:]
		} else if donti < doi && donti < muli {
			iter = append(iter, dont{})
			donts = donts[1:]
		} else {
			mul := muls[0]
			l := d.parseInt(input[mul[2]:mul[3]])
			r := d.parseInt(input[mul[4]:mul[5]])
			iter = append(iter, multiply{l: l, r: r})
			muls = muls[1:]
		}

		empty = len(dos) == 0 && len(donts) == 0 && len(muls) == 0
		if empty {
			break
		}
	}
	return iter
}

func (d *Day3) printIter(iter []any) {
	result := 0
	enabled := true
	for _, item := range iter {
		switch i := item.(type) {
		case do:
			enabled = true
			// fmt.Println(Reset + "do")
		case dont:
			enabled = false
			// fmt.Println(Reset + "dont")
		case multiply:
			color := Green
			mul := i.l * i.r
			if enabled {
				result += mul
				fmt.Printf(color+"%v,%v,%v,%v\n", i.l, i.r, mul, result)
			}
		}
	}
}

func (d *Day3) getIndexes(matches [][]int) []int {
	indexes := []int{}
	for _, match := range matches {
		indexes = append(indexes, match[0])
	}
	return indexes
}

func (d *Day3) parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		util.LogErr(err)
		// return -1 to make error more obvious
		return -1
	}
	return i
}
