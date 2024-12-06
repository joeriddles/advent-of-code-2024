package main

import (
	"testing"

	"github.com/joeriddles/advent-of-code-2024/pkg/util"
)

func TestPart1_Horizontal(t *testing.T) {
	day := &Day4{}

	input := `.XMAS.`
	util.Assert(t, day.Part1(input), 1)

	input = `....XXMAS.`
	util.Assert(t, day.Part1(input), 1)
}

func TestPart1_Vertical(t *testing.T) {
	day := &Day4{}

	input := `.
X
M
A
S
.`
	if actual := day.Part1(input); actual != 1 {
		t.Fatalf("expected %v, got %v", 1, actual)
	}
}

func TestPart1_Backwards(t *testing.T) {
	day := &Day4{}
	input := `.SAMX.`
	util.Assert(t, day.Part1(input), 1)

	input = `XMSMSAMASM`
	util.Assert(t, day.Part1(input), 0)
}

func TestPart1_DiagonalSimple(t *testing.T) {
	day := &Day4{}
	input := `
X...
.M..
..A.
...S`
	util.Assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalBackwards(t *testing.T) {
	day := &Day4{}
	input := `
...X
..M.
.A..
S...`
	util.Assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalExtraLine(t *testing.T) {
	day := &Day4{}
	input := `....
X...
.M..
..A.
...S`
	util.Assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalUpsideDown(t *testing.T) {
	day := &Day4{}
	input := `
...S
..A.
.M..
X...`
	util.Assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalUpsideDownAndBackwards(t *testing.T) {
	day := &Day4{}
	input := `
S...
.A..
..M.
...X`
	util.Assert(t, day.Part1(input), 1)
}

func TestPart1_Placeholders(t *testing.T) {
	day := &Day4{}

	input := `
..X...
.SAMX.
.A..A.
XMAS.S
.X....`
	util.Assert(t, day.Part1(input), 4)
	// horizontal = 1
	// horizontalBackwards = 1
	// verticalUpsidedown = 1
	// diagonal = 1

	input = `
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM`
	util.Assert(t, day.Part1(input), 6)
	// horizontal = 2
	// horizontalBackwards = 2
	// verticalUpsidedown = 1
	// diagonal = 1

	input = `
....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A`
	util.Assert(t, day.Part1(input), 8)

	input = `....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`
	util.Assert(t, day.Part1(input), 18)
}

func TestPart1(t *testing.T) {
	day := &Day4{}

	// ....XXMAS. <- h
	// .SAMXMS... <- hb
	// ...S..A...
	// ..A.A.MS.X
	// XMASAMX.MM <- h, hb
	input := `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM`
	util.Assert(t, day.Part1(input), 6)
	// horizontal = 2
	// horizontalBackwards = 2
	// verticalUpsidedown = 1
	// diagonal = 1

	input = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA`
	util.Assert(t, day.Part1(input), 8)

	input = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	util.Assert(t, day.Part1(input), 18)
}

func TestPart2(t *testing.T) {
	day := &Day4{}

	src := `
M.S
.A.
M.S`
	util.Assert(t, day.Part2(src), 1)

	src = `
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.`
	util.Assert(t, day.Part2(src), 4)

	src = `
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`
	util.Assert(t, day.Part2(src), 9)
}
