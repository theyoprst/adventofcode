package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	ans := 0
	for _, s := range lines {
		ans += s2i(s)
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	ans := 0
	for _, s := range lines {
		for i, word := range []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"} {
			s = strings.ReplaceAll(s, word, fmt.Sprintf("%s%d%s", word, i+1, word))
		}
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
