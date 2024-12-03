package main

import "testing"

const src = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	expected := 2
	actual := (&Day2{}).Part1(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 4
	actual := (&Day2{}).Part2(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
