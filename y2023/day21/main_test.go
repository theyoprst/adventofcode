package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "42",
		},
		{
			Path:      "input.txt",
			WantPart1: "3632",
			WantPart2: "600336060511101",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
