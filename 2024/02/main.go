package main

import (
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(lines []string) any {
	safeCount := 0
	for _, report := range lines {
		levels := aoc.Ints(report)
		if isSafe(levels) {
			safeCount++
		}
	}
	return safeCount
}

func SolvePart2(lines []string) any {
	safeCount := 0
	for _, report := range lines {
		levels := aoc.Ints(report)
		if isAlmostSafe(levels) {
			safeCount++
		}
	}
	return safeCount
}

func isAlmostSafe(levels []int) bool {
	if IsAlmostSafeIncreasing(levels) {
		return true
	}
	slices.Reverse(levels)
	return IsAlmostSafeIncreasing(levels)
}

func isSafe(levels []int) bool {
	if IsSafeIncreasing(levels) {
		return true
	}
	slices.Reverse(levels)
	return IsSafeIncreasing(levels)
}

func IsSafeIncreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		inc := levels[i] - levels[i-1]
		if inc < 1 || inc > 3 {
			return false
		}
	}
	return true
}

func isSafeInc(lo, hi int) bool {
	return hi-lo >= 1 && hi-lo <= 3
}

func IsAlmostSafeIncreasing(levels []int) bool {
	var breaks []int
	for i := 1; i < len(levels); i++ {
		if !isSafeInc(levels[i-1], levels[i]) {
			breaks = append(breaks, i)
		}
	}
	if len(breaks) == 0 {
		return true
	}
	if len(breaks) == 1 {
		j := breaks[0]
		// Try to remove j-1
		if j == 1 {
			return true
		}
		if isSafeInc(levels[j-2], levels[j]) {
			return true
		}
		// Try to remove j
		if j == len(levels)-1 {
			return true
		}
		return isSafeInc(levels[j-1], levels[j+1])
	}
	if len(breaks) == 2 {
		if breaks[1] != breaks[0]+1 {
			return false
		}
		j := breaks[0]
		// Try to remove j
		return isSafeInc(levels[j-1], levels[j+1])
	}
	return false
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
