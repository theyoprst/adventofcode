package main

import (
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	return slices.Max(blockSums(lines))
}

func SolvePart2(lines []string) any {
	sums := blockSums(lines)
	slices.Sort(sums)
	var ans int
	for _, s := range sums[len(sums)-3:] {
		ans += s
	}
	return ans
}

func blockSums(lines []string) []int {
	var sums []int
	for _, block := range aoc.Split(lines, "") {
		sum := 0
		for _, line := range block {
			sum += must.Atoi(line)
		}
		sums = append(sums, sum)
	}
	return sums
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
