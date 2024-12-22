package main

import (
	"context"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func Next(a []int) int {
	a = aoc.Reversed(a)
	next := 0
	for size := len(a); size >= 2; size-- {
		next += a[0]
		for i := 0; i < size-1; i++ {
			a[i] -= a[i+1]
		}
	}
	must.Equal(a[0], 0)
	return next
}

func SolvePart1(_ context.Context, lines []string) any {
	var ans int
	for _, line := range lines {
		ans += Next(aoc.Ints(line))
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	var ans int
	for _, line := range lines {
		a := aoc.Ints(line)
		ans += Next(aoc.Reversed(a))
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
