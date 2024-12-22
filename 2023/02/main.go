package main

import (
	"context"
	"log"
	"regexp"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	var ans int
	for gameI, line := range lines {
		c := parseColors(line)
		if c["red"] <= 12 && c["green"] <= 13 && c["blue"] <= 14 {
			ans += gameI + 1
		}
	}
	return ans
}

func SolvePart2(_ context.Context, lines []string) any {
	var ans int
	for _, line := range lines {
		c := parseColors(line)
		ans += c["red"] * c["green"] * c["blue"]
	}
	return ans
}

var reColor = regexp.MustCompile(` (\d+) +(\w+)`)

func parseColors(game string) map[string]int {
	maxColors := map[string]int{}
	for _, match := range reColor.FindAllStringSubmatch(game, -1) {
		color := match[2]
		maxColors[color] = max(maxColors[color], must.Atoi(match[1]))
	}
	return maxColors
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	log.SetFlags(0)
	aoc.Main(solvers1, solvers2)
}
