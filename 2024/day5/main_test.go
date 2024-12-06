package main

import (
	"os"
	"testing"

	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

const src = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func Test_validate(t *testing.T) {
	actual := validate(7, 4, map[int][]int{
		4: {9},
		9: {7},
	})
	if actual {
		t.Fatal("expected not valid")
	}
}

func TestPart1_Simple(t *testing.T) {
	input := `
1|2
2|3

1,2,3`
	actual := (&Day5{}).Part1(input)
	util.Assert(t, actual, 2)
}

func TestPart1_SimpleBackwards(t *testing.T) {
	input := `
3|2
2|1

3,2,1`
	actual := (&Day5{}).Part1(input)
	util.Assert(t, actual, 2)
}

func TestPart1(t *testing.T) {
	actual := (&Day5{}).Part1(src)
	util.Assert(t, actual, 143)
}

func TestPart2(t *testing.T) {
	expected := 0
	actual := (&Day5{}).Part2(src)
	util.Assert(t, actual, expected)
}

func TestPart1_Real(t *testing.T) {
	debug = true
	bytes, _ := os.ReadFile("./input.txt")
	input := string(bytes)
	(&Day5{}).Part1(input)
}
