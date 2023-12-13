package main

import (
	"github.com/theyoprst/adventofcode/aoc"
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
		field := aoc.ToBytesField(pattern)
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

var solversPart1 []aoc.Solver = []aoc.Solver{
	SolvePart1,
}

var solversPart2 []aoc.Solver = []aoc.Solver{
	SolvePart2,
}

func main() {
	aoc.Main([]aoc.Solver{SolvePart1}, []aoc.Solver{SolvePart2})
}
