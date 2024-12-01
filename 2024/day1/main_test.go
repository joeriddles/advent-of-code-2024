package main

import "testing"

func TestMain(t *testing.T) {
	src := `3   4
4   3
2   5
1   3
3   9
3   3
`

	expected := 11
	actual := test(src)

	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
