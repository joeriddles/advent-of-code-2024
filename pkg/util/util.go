package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(i int) int {
	if i < 0 {
		i = -1 * i
	}
	return i
}

// Parse newline and whitespace separated numebrs into a slice of int slices
func ParseIntSlices(input string) []*[]int {
	reports := []*[]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		report := []int{}
		for _, part := range parts {
			part = strings.Trim(part, " ")
			level, _ := strconv.Atoi(part)
			report = append(report, level)
		}
		reports = append(reports, &report)
	}
	return reports
}

func HeadOrDefault[T any](slice []T, def T) T {
	if len(slice) == 0 {
		return def
	}
	return slice[0]
}

func LogErr(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
}
