package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "5",
			WantPart2: "7",
		},
		{
			Path:      "input.txt",
			WantPart1: "430",
			WantPart2: "60558",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
