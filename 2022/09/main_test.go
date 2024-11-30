package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "13",
			WantPart2: "1",
		},
		{
			Path:      "input_ex2.txt",
			WantPart1: "88",
			WantPart2: "36",
		},
		{
			Path:      "input.txt",
			WantPart1: "6339",
			WantPart2: "2541",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
