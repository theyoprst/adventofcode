package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "6440",
			WantPart2: "5905",
		},
		{
			Path:      "input.txt",
			WantPart1: "246424613",
			WantPart2: "248256639",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
