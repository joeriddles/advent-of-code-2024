package util

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func Abs(i int) int {
	if i < 0 {
		i = -1 * i
	}
	return i
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("%v is not an int", s))
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

const (
	Reset string = "\033[0m"
	Red   string = "\033[31m"
	Green string = "\033[32m"
)

func LogErr(err error) {
	fmt.Fprintf(os.Stderr, Red+"%s\n"+Reset, err.Error())
}

func LogSuccessf(format string, args ...any) {
	fmt.Printf(Green+format+Reset, args...)
}

func LogErrf(format string, args ...any) {
	fmt.Printf(Red+format+Reset, args...)
}

func Map[T any, R any](s []T, f func(T) R) []R {
	rs := []R{}
	for _, t := range s {
		rs = append(rs, f(t))
	}
	return rs
}

func Where[T any](s []T, f func(T) bool) []T {
	res := []T{}
	for _, t := range s {
		if f(t) {
			res = append(res, t)
		}
	}
	return res
}

func Assert(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func IsDebug() bool {
	return slices.Contains(os.Args, "--debug")
}
