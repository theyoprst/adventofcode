package main

import (
	"math"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(lines []string) any {
	times := aoc.Ints(lines[0])
	records := aoc.Ints(lines[1])
	return CountWins(times, records)
}

func SolvePart2(lines []string) any {
	times := []int{ParseSplittedInt(lines[0])}
	records := []int{ParseSplittedInt(lines[1])}
	return CountWins(times, records)
}

func SolvePart1Fast(lines []string) any {
	times := aoc.Ints(lines[0])
	records := aoc.Ints(lines[1])
	return CountWinsFast(times, records)
}

func SolvePart2Fast(lines []string) any {
	times := []int{ParseSplittedInt(lines[0])}
	records := []int{ParseSplittedInt(lines[1])}
	return CountWinsFast(times, records)
}

func CountWins(times, records []int) int {
	ans := 1
	for i := range times {
		t := times[i]
		r := records[i]
		wins := 0
		for x := 0; x <= t; x++ {
			y := times[i] - x
			wins += aoc.BoolToInt(x*y > r)
		}
		ans *= wins
	}
	return ans
}

func CountWinsFast(times, records []int) int {
	ans := 1
	for i := range times {
		t := times[i]
		r := records[i]
		// x * (t - x) > r <=> -x*x +tx -r > 0
		x1, x2 := aoc.SolveQuadratic(-1, t, -r)
		const eps = 1e-10
		p1 := int(math.Ceil(*x1 + eps))
		p2 := int(math.Floor(*x2 - eps))
		ans *= p2 - p1 + 1
	}
	return ans
}

func ParseSplittedInt(s string) int {
	n := 0
	for _, ch := range s {
		if aoc.IsDigit(ch) {
			n *= 10
			n += int(ch - '0')
		}
	}
	return n
}

var (
	solvers1 = []aoc.Solver{SolvePart1, SolvePart1Fast}
	solvers2 = []aoc.Solver{SolvePart2, SolvePart2Fast}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
