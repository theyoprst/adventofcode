package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(_ context.Context, lines []string) any {
	_ = lines
	return 0
}

func SolvePart2(_ context.Context, lines []string) any {
	_ = lines
	return 0
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
