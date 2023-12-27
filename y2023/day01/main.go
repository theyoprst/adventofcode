package main

import (
	"log"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	ans := 0
	for _, s := range lines {
		ans += s2i(s)
	}
	return ans
}

func SolvePart2(lines []string) any {
	ans := 0
	for _, s := range lines {
		s = strings.ReplaceAll(s, "one", "o1e")
		s = strings.ReplaceAll(s, "two", "t2o")
		s = strings.ReplaceAll(s, "three", "t3e")
		s = strings.ReplaceAll(s, "four", "f4r")
		s = strings.ReplaceAll(s, "five", "f5e")
		s = strings.ReplaceAll(s, "six", "s6x")
		s = strings.ReplaceAll(s, "seven", "s7n")
		s = strings.ReplaceAll(s, "eight", "e8t")
		s = strings.ReplaceAll(s, "nine", "n9e")
		ans += s2i(s)
	}
	return ans
}

func s2i(s string) int {
	first, last := -1, -1
	for _, r := range s {
		if aoc.IsDigit(r) {
			if first == -1 {
				first = int(r - '0')
			}
			last = int(r - '0')
		}
	}
	must.NotEqual(first, -1)
	return first*10 + last
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	log.SetFlags(0)
	aoc.Main(solvers1, solvers2)
}
