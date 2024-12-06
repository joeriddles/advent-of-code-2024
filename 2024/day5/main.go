package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/joeriddles/advent-of-code-2024/pkg/day"
	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

var debug = false

const MaxInt = int(^uint(0) >> 1)

func main() {
	debug = util.IsDebug()

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

type node struct {
	val   int
	nodes []*node
}

func newNode(val int) *node {
	return &node{val: val, nodes: []*node{}}
}

func (n *node) valid(r int) bool {
	for _, node := range n.nodes {
		if node.val == r {
			return true
		}
		if node.valid(r) {
			return true
		}
	}
	return false
}

func (d *Day5) Part1(input string) int {
	rules, updates := d.parse(input)
	_ = updates

	nodes := map[int]*node{}
	for _, rule := range rules {
		l, r := rule[0], rule[1]

		if _, ok := nodes[l]; !ok {
			nodes[l] = newNode(l)
		}
		if _, ok := nodes[r]; !ok {
			nodes[r] = newNode(r)
		}

		nodes[l].nodes = append(nodes[l].nodes, nodes[r])
	}

	for _, update := range updates {
		l := update[0]
		valid := true
		for _, r := range update[1:] {
			if !nodes[l].valid(r) {
				valid = false
				break
			}
			l = r
		}
		fmt.Printf("%v: %v\n", update, valid)
	}

	// Build ruleMap of each page and what pages it must come before
	ruleMap := map[int][]int{}
	for _, rule := range rules {
		l, r := rule[0], rule[1]
		ruleMap[l] = append(ruleMap[l], r)

		// Make sure the page last value is added, too
		if _, ok := ruleMap[r]; !ok {
			ruleMap[r] = []int{}
		}
	}

	allPages := make([]int, 0, len(ruleMap))
	for p := range ruleMap {
		allPages = append(allPages, p)
	}

	// Filter out nonsense pages
	for p, rules := range ruleMap {
		ruleMap[p] = util.Where(rules, func(r int) bool { return slices.Contains(allPages, p) })
	}

	order := []int{}
	for {
		if len(order) == len(allPages) {
			break
		}

		for page, rules := range ruleMap {
			for _, rule := range rules {
				if !slices.Contains(allPages, rule) {
					fmt.Printf("%v -> %v\n", page, rule)
				}
			}

			if len(rules) == 0 {
				order = slices.Insert(order, 0, page)
				// prune page from rules
				for p, rules := range ruleMap {
					ruleMap[p] = util.Where(rules, func(p int) bool { return p != page })
				}
				delete(ruleMap, page)
				break
			}
		}
	}

	// result := 0
	// NOTE: this stack overflows, super inefficient...
	// for _, update := range updates {
	// 	l := update[0]
	// 	valid := true
	// 	for _, r := range update[1:] {
	// 		if !validate_cache(l, r, ruleMap) {
	// 			valid = false
	// 			break
	// 		}
	// 		l = r
	// 	}
	// 	if valid {
	// 		mid := update[len(update)/2]
	// 		result += mid
	// 	}
	// }

	// // Order page rules
	// for p, lts := range lookup {
	// 	slices.Sort(lts)

	// 	// Add a rule for any pages that are missing a rule
	// 	if len(lts) == 0 {
	// 		lookup[p] = append(lts, MaxInt)
	// 	}
	// }

	// // Build correct order of pages by continually grabbing the page with the lowest page rule
	// NOTE: this doesn't work correctly...
	// order := []int{}
	// for {
	// 	if len(lookup) == 0 {
	// 		break
	// 	}

	// 	minPage := -1
	// 	min := -1
	// 	for page, lts := range lookup {
	// 		lt := lts[0]
	// 		if min == -1 || lt < min {
	// 			min = lt
	// 			minPage = page
	// 			continue
	// 		}
	// 	}

	// 	delete(lookup, minPage)
	// 	order = append(order, minPage)
	// }

	// Check ordering of updates
	result := 0
	for _, update := range updates {
		cur := -1
		ordered := true
		for _, page := range update {
			i := slices.Index(order, page)
			if i < cur {
				ordered = false
				break
			}
			cur = i
		}
		if ordered {
			mid := update[len(update)/2]
			result += mid
		}
	}

	return result
}

func (d *Day5) Part2(input string) int {
	return -1
}

func (d *Day5) parse(input string) ([][]int, [][]int) {
	lines := strings.Split(input, "\n")
	rules := [][]int{}
	updates := [][]int{}
	for _, line := range lines {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			rule := util.Map(parts, util.MustParseInt)
			rules = append(rules, rule)
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			update := util.Map(parts, util.MustParseInt)
			updates = append(updates, update)
		}
	}
	return rules, updates
}

type pair struct {
	l, r int
}

func (p pair) String() string {
	return fmt.Sprintf("%v|%v", p.l, p.r)
}

var lru = map[pair]bool{}

func validate_cache(l, r int, rules map[int][]int) bool {
	p := pair{l: l, r: r}
	if valid, ok := lru[p]; ok {
		return valid
	}
	valid := validate(l, r, rules)
	if debug {
		fmt.Printf("%v: %v\n", p, valid)
	}
	lru[p] = valid
	return valid
}

// Validate l is less than r based on the rules
func validate(l, r int, rules map[int][]int) bool {
	rrules := rules[r]
	for _, lt := range rrules {
		if lt == l {
			return false
		}
		if !validate(l, lt, rules) {
			return false
		}
	}
	return true
}
