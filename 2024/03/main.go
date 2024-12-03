package main

import (
	"strings"

	"github.com/hedhyw/rex/pkg/rex"
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
	// Traditional way to declare the regex, not quite readable:
	// re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Functional way to declare the regex: too verbose, but more readable:
	// re := rex.New(
	// 	rex.Common.Raw(`mul`),
	// 	rex.Chars.Single('('),
	// 	rex.Group.Define(rex.Chars.Digits().Repeat().OneOrMore()),
	// 	rex.Chars.Single(','),
	// 	rex.Group.Define(rex.Chars.Digits().Repeat().OneOrMore()),
	// 	rex.Chars.Single(')'),
	// ).MustCompile()

	// Best of two worlds: same regex with extra spaces and comments.
	re := rex.New(rex.Common.RawVerbose(`
		mul\(      # Beware: spaces are trimmed in RawVerbose, but spaces in the middle are retained
		           # (a space between "mul" and "\(" will break the regex).
			(\d+), # Group 1: first number.
			(\d+)  # Group 2: second number.
		\)
	`)).MustCompile()

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
