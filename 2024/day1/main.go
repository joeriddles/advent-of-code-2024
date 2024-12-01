package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("single filepath required")
		os.Exit(1)
	}
	fp := os.Args[1]
	bytes, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	input := string(bytes)

	answer1 := part1(input)
	fmt.Printf("answer 1: %v\n", answer1)

	answer2 := part2(input)
	fmt.Printf("answer 2: %v\n", answer2)
}

func part1(src string) int {
	lists := parse(src)
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

func part2(src string) int {
	lists := parse(src)
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

func parse(src string) []*[]int {
	lines := strings.Split(src, "\n")
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
