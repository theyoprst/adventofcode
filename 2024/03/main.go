package main

import (
	"regexp"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(lines []string) any {
	return countMulSum(strings.Join(lines, "\n"))
}

func SolvePart2(lines []string) any {
	edited := removeDisabled(strings.Join(lines, "\n"))
	return countMulSum(edited)
}

func removeDisabled(text string) string {
	const (
		do   = "do()"
		dont = "don't()"
	)
	var (
		enable = true
		edited = make([]byte, 0, len(text))
	)
	for i := range text {
		if strings.HasPrefix(text[i:], do) {
			enable = true
		} else if strings.HasPrefix(text[i:], dont) {
			enable = false
		}
		if enable {
			edited = append(edited, text[i])
		}
	}
	return string(edited)
}

func countMulSum(text string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	sum := 0
	for _, match := range re.FindAllStringSubmatch(text, -1) {
		sum += must.Atoi(match[1]) * must.Atoi(match[2])
	}
	return sum
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
