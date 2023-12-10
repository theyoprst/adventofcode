package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{Path: "input_ex0.txt", WantPart1: "4", WantPart2: "1"},
		{Path: "input_ex1.txt", WantPart1: "4", WantPart2: "1"},
		{Path: "input_ex2.txt", WantPart1: "8", WantPart2: "1"},
		{Path: "input_ex3.txt", WantPart2: "4"},
		{Path: "input_ex4.txt", WantPart2: "8"},
		{Path: "input.txt", WantPart1: "6897", WantPart2: "367"},
	}
	solversPart1 := []aoc.Solver{
		SolvePart1,
	}
	solversPart2 := []aoc.Solver{
		SolvePart2,
	}
	aoc.RunTests(t, inputs, solversPart1, solversPart2)
}
