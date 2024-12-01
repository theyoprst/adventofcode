package main

import (
	"sort"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(lines []string) any {
	var x, y []int
	for _, line := range lines {
		ii := aoc.Ints(line)
		x = append(x, ii[0])
		y = append(y, ii[1])
	}
	sort.Ints(x)
	sort.Ints(y)
	diff := 0
	for i := range x {
		diff += aoc.Abs(x[i] - y[i])
	}
	return diff
}

func SolvePart2(lines []string) any {
	var x []int
	counter := map[int]int{}
	for _, line := range lines {
		ii := aoc.Ints(line)
		x = append(x, ii[0])
		counter[ii[1]]++
	}
	score := 0
	for i := range x {
		score += counter[x[i]] * x[i]
	}
	return score
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
