package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	var ans int
	for _, line := range lines {
		a, b := must.Split2(line, ",")
		a1Str, a2Str := must.Split2(a, "-")
		b1Str, b2Str := must.Split2(b, "-")
		a1 := must.Atoi(a1Str)
		a2 := must.Atoi(a2Str)
		b1 := must.Atoi(b1Str)
		b2 := must.Atoi(b2Str)
		must.LessOrEqual(a1, a2)
		must.LessOrEqual(b1, b2)
		// Intersect:
		i1 := max(a1, b1)
		i2 := min(a2, b2)
		if i1 == a1 && i2 == a2 || i1 == b1 && i2 == b2 {
			ans++
		}
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	var ans int
	for _, line := range lines {
		a, b := must.Split2(line, ",")
		a1Str, a2Str := must.Split2(a, "-")
		b1Str, b2Str := must.Split2(b, "-")
		a1 := must.Atoi(a1Str)
		a2 := must.Atoi(a2Str)
		b1 := must.Atoi(b1Str)
		b2 := must.Atoi(b2Str)
		must.LessOrEqual(a1, a2)
		must.LessOrEqual(b1, b2)
		// Intersect:
		i1 := max(a1, b1)
		i2 := min(a2, b2)
		if i1 <= i2 {
			ans++
		}
	}
	return ans
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
