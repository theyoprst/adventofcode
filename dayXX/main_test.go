package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{Path: "input_ex1.txt", WantPart1: "0", WantPart2: "0"},
		{Path: "input.txt", WantPart1: "0", WantPart2: "0"},
	}
	solversPart1 := []aoc.Solver{
		SolvePart1,
	}
	solversPart2 := []aoc.Solver{
		SolvePart2,
	}
	aoc.RunTests(t, inputs, solversPart1, solversPart2)
}
