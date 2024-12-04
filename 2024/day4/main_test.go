package main

import "testing"

func TestPart1_Horizontal(t *testing.T) {
	day := &Day4{}

	input := `.XMAS.`
	if actual := day.Part1(input); actual != 1 {
		t.Fatalf("expected %v, got %v", 1, actual)
	}
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
	assert(t, day.Part1(input), 1)

	input = `XMSMSAMASM`
	assert(t, day.Part1(input), 0)
}

func TestPart1_DiagonalSimple(t *testing.T) {
	day := &Day4{}
	input := `
X...
.M..
..A.
...S`
	assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalBackwards(t *testing.T) {
	day := &Day4{}
	input := `
...X
..M.
.A..
S...`
	assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalExtraLine(t *testing.T) {
	day := &Day4{}
	input := `....
X...
.M..
..A.
...S`
	assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalUpsideDown(t *testing.T) {
	day := &Day4{}
	input := `
...S
..A.
.M..
X...`
	assert(t, day.Part1(input), 1)
}

func TestPart1_DiagonalUpsideDownAndBackwards(t *testing.T) {
	day := &Day4{}
	input := `
S...
.A..
..M.
...X`
	assert(t, day.Part1(input), 1)
}

func TestPart1_Placeholders(t *testing.T) {
	day := &Day4{}

	input := `
..X...
.SAMX.
.A..A.
XMAS.S
.X....`
	assert(t, day.Part1(input), 4)
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
	assert(t, day.Part1(input), 6)
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
	assert(t, day.Part1(input), 8)

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
	assert(t, day.Part1(input), 18)
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
	assert(t, day.Part1(input), 6)
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
	assert(t, day.Part1(input), 8)

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
	assert(t, day.Part1(input), 18)
}

// func TestPart2(t *testing.T) {
// 	expected := 0
// 	actual := (&Day4{}).Part2(src)
// 	if actual != expected {
// 		t.Fatalf("expected %v, got %v", expected, actual)
// 	}
// }

func assert(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
