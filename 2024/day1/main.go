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
	answer := test(input)
	fmt.Printf("answer: %v\n", answer)
}

func test(src string) int {
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

func abs(i int) int {
	if i < 0 {
		i = -1 * i
	}
	return i
}
