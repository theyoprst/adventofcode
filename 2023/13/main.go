package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

func FindMirrorPoint(a []int, wantMismatches int) int {
	for i := 1; i < len(a); i++ {
		// Count a[:i] and a[i:] mismatches.
		left, right := i-1, i
		mismatches := 0
		for 0 <= left && right < len(a) { // && mismatches <= wantMismatches could speed it up a bit (but it doesn't on these data)
			mismatches += aoc.CountBinaryOnes(a[left] ^ a[right])
			left--
			right++
		}
		if mismatches == wantMismatches {
			return i
		}
	}
	return 0
}

func SolveGeneric(lines []string, wantMismatches int) any {
	var ans int
	for _, pattern := range aoc.Split(lines, "") {
		field := fld.NewByteField(pattern)
		rowMasks := make([]int, len(field))
		colMasks := make([]int, len(field[0]))
		for row, line := range field {
			for col, ch := range line {
				rowMasks[row] *= 2
				colMasks[col] *= 2
				if ch == '#' {
					rowMasks[row]++
					colMasks[col]++
				}
			}
		}
		ans += FindMirrorPoint(colMasks, wantMismatches) + 100*FindMirrorPoint(rowMasks, wantMismatches)
	}
	return ans
}

func SolvePart1(lines []string) any {
	return SolveGeneric(lines, 0)
}

func SolvePart2(lines []string) any {
	return SolveGeneric(lines, 1)
}

//
// Next goes solution without integer masks, just matrices.
//

func Mismatches(a, b []byte) int {
	must.Equal(len(a), len(b))
	mismatches := 0
	for i := range a {
		mismatches += aoc.BoolToInt(a[i] != b[i])
	}
	return mismatches
}

func HorMirrorPoint(field [][]byte, wantMismatches int) int {
	for row := 1; row < len(field); row++ {
		top, bottom := row-1, row
		diff := 0
		for 0 <= top && bottom < len(field) {
			diff += Mismatches(field[top], field[bottom])
			top--
			bottom++
		}
		if diff == wantMismatches {
			return row
		}
	}
	return 0
}

func SolvePart1Transponse(lines []string) any {
	var ans int
	for _, pattern := range aoc.Split(lines, "") {
		field := fld.NewByteField(pattern)
		trans := field.NewFieldTransposed()
		ans += 100*HorMirrorPoint(field, 0) + HorMirrorPoint(trans, 0)
	}
	return ans
}

func SolvePart2Transponse(lines []string) any {
	var ans int
	for _, pattern := range aoc.Split(lines, "") {
		field := fld.NewByteField(pattern)
		trans := field.NewFieldTransposed()
		ans += 100*HorMirrorPoint(field, 1) + HorMirrorPoint(trans, 1)
	}
	return ans
}

var solversPart1 = []aoc.Solver{
	SolvePart1,
	SolvePart1Transponse,
}

var solversPart2 = []aoc.Solver{
	SolvePart2,
	SolvePart2Transponse,
}

func main() {
	aoc.Main([]aoc.Solver{SolvePart1}, []aoc.Solver{SolvePart2})
}
