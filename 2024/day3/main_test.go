package main

import "testing"

const src = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestPart1(t *testing.T) {
	expected := 161
	actual := (&Day3{}).Part1(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 48
	actual := (&Day3{}).Part2(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
