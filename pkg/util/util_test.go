package util_test

import (
	"slices"
	"testing"

	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

func TestParseIntSlices_Simple(t *testing.T) {
	actual := util.ParseIntSlices(`3 4
4 3`)
	expected := []*[]int{
		{3, 4},
		{4, 3},
	}

	for i, exp := range expected {
		act := actual[i]
		if !slices.Equal(*exp, *act) {
			t.Fatalf("expected %v, got %v", exp, act)
		}
	}
}

func TestParseIntSlices_ExtraSpaces(t *testing.T) {
	actual := util.ParseIntSlices(`3   4
4   3`)
	expected := []*[]int{
		{3, 4},
		{4, 3},
	}

	for i, exp := range expected {
		act := actual[i]
		if !slices.Equal(*exp, *act) {
			t.Fatalf("expected %v, got %v", exp, act)
		}
	}
}
