package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
)

func SolvePart1(lines []string) any {
	return solve(lines, 4)
}

func SolvePart2(lines []string) any {
	return solve(lines, 14)
}

func solve(lines []string, width int) int {
	for _, line := range lines {
		for i := width; i < len(line); i++ {
			sub := []byte(line[i-width : i])
			if len(containers.NewSet[byte](sub...)) == len(sub) {
				return i
			}
		}
	}
	panic("unreachable")
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
