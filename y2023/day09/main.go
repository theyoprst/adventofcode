package main

import (
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

func SolvePart1(lines []string) any {
	var ans int
	for _, line := range lines {
		ans += Next(aoc.Ints(line))
	}
	return ans
}

func SolvePart2(lines []string) any {
	var ans int
	for _, line := range lines {
		a := aoc.Ints(line)
		ans += Next(aoc.Reversed(a))
	}
	return ans
}

var (
	solvers1 []aoc.Solver = []aoc.Solver{SolvePart1}
	solvers2 []aoc.Solver = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
