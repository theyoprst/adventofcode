package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "136",
			WantPart2: "64",
		},
		{
			Path:      "input.txt",
			WantPart1: "108955",
			WantPart2: "106689",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
