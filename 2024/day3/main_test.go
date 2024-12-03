package main

import "testing"

const src = ``

func TestPart1(t *testing.T) {
	expected := 0
	actual := (&Day3{}).Part1(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	actual := (&Day3{}).Part2(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
