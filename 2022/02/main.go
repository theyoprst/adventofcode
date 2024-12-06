// 14:00 - 14:17 - 14:27.
package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/must"
)

type Figure int

const (
	Rock Figure = iota
	Paper
	Scissors
)

type Outcome string

const (
	Lose = "X"
	Draw = "Y"
	Win  = "Z"
)

func norm(c string) Figure {
	switch c {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		panic("Unreachable")
	}
}

func getOutcome(f1, f2 Figure) Outcome {
	if f1 == f2 {
		return Draw
	}
	if next[f1] == f2 {
		return Win
	}
	return Lose
}

var next = map[Figure]Figure{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var figureScore = map[Figure]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var outcomeScore = map[Outcome]int{
	Lose: 0,
	Draw: 3,
	Win:  6,
}

func score(f1, f2 Figure) int {
	return figureScore[f2] + outcomeScore[getOutcome(f1, f2)]
}

func score2(f Figure, outcome Outcome) int {
	for _, f2 := range []Figure{Rock, Paper, Scissors} {
		if getOutcome(f, f2) == outcome {
			return score(f, f2)
		}
	}
	panic("Unreachable")
}

func SolvePart1(lines []string) any {
	var ans int
	for _, line := range lines {
		c1, c2 := must.Split2(line, " ")
		ans += score(norm(c1), norm(c2))
	}
	return ans
}

func SolvePart2(lines []string) any {
	var ans int
	for _, line := range lines {
		c1, c2 := must.Split2(line, " ")
		ans += score2(norm(c1), Outcome(c2))
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
