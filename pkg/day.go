package pkg

import (
	"fmt"
	"os"
)

type Day interface {
	Part1(string) int
	Part2(string) int
}

type BaseDay struct{}

func NewBaseDay() *BaseDay {
	return &BaseDay{}
}

func (*BaseDay) Parse() string {
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
	return input
}

// Part1 implements Day.
func (b *BaseDay) Part1(string) int {
	panic("unimplemented")
}

// Part2 implements Day.
func (b *BaseDay) Part2(string) int {
	panic("unimplemented")
}
