package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg"
)

func main() {
	day1 := &Day1{}
	input := day1.Parse()

	answer1 := day1.Part1(input)
	fmt.Printf("part1: %v\n", answer1)

	answer2 := day1.Part2(input)
	fmt.Printf("part2: %v\n", answer2)
}

type Day1 struct {
	*pkg.BaseDay
}

func NewDay1() pkg.Day {
	return &Day1{
		BaseDay: pkg.NewBaseDay(),
	}
}

func (d *Day1) Part1(src string) int {
	lists := d.parseInts(src)
	ls := *lists[0]
	rs := *lists[1]
	sort.Ints(ls)
	sort.Ints(rs)

	res := 0
	for i, l := range ls {
		r := rs[i]
		diff := abs(r - l)
		res += diff
	}

	return res
}

func (d *Day1) Part2(src string) int {
	lists := d.parseInts(src)
	ls := *lists[0]
	rs := *lists[1]

	rmap := map[int]int{}
	for _, r := range rs {
		if _, ok := rmap[r]; !ok {
			rmap[r] = 0
		}
		rmap[r]++
	}

	res := 0
	for _, l := range ls {
		rcount, ok := rmap[l]
		if !ok {
			rcount = 0
		}

		sim := l * rcount
		res += sim
	}

	return res
}

func (d *Day1) parseInts(input string) []*[]int {
	lines := strings.Split(input, "\n")
	ls := []int{}
	rs := []int{}
	for _, line := range lines {
		if line == "" {
			break
		}

		parts := strings.SplitN(line, "   ", 2)
		lstr := parts[0]
		rstr := parts[1]
		l, _ := strconv.Atoi(lstr)
		r, _ := strconv.Atoi(rstr)
		ls = append(ls, l)
		rs = append(rs, r)
	}

	return []*[]int{&ls, &rs}
}

func abs(i int) int {
	if i < 0 {
		i = -1 * i
	}
	return i
}
