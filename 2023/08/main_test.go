package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "2",
			WantPart2: "2",
		},
		{
			Path:      "input_ex2.txt",
			WantPart1: "6",
		},
		{
			Path:      "input_ex3.txt",
			WantPart2: "6",
		},
		{
			Path:      "input.txt",
			WantPart1: "12643",
			WantPart2: "13133452426987",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
