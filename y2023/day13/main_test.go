package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{Path: "input_ex1.txt", WantPart1: "405", WantPart2: "400"},
		{Path: "input.txt", WantPart1: "37025", WantPart2: "32854"},
	}
	aoc.RunTests(t, inputs, solversPart1, solversPart2)
}
