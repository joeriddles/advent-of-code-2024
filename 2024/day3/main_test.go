package main

import "testing"

func TestPart1(t *testing.T) {
	src := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	expected := 161
	actual := (&Day3{}).Part1(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	src := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := 48
	actual := (&Day3{}).Part2(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestPart2_Custom(t *testing.T) {
	src := `mul(1,1)don't()mul(1,2)do()mul(1,3)`
	expected := 4
	actual := (&Day3{}).Part2(src)
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
