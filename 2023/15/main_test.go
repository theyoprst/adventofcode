package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "52",
		},
		{
			Path:      "input_ex2.txt",
			WantPart1: "1320",
			WantPart2: "145",
		},
		{
			Path:      "input.txt",
			WantPart1: "515974",
			WantPart2: "265894",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
