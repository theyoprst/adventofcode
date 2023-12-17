package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "102",
			WantPart2: "94",
		},
		{
			Path:      "input_ex2.txt",
			WantPart1: "59",
			WantPart2: "47",
		},
		{
			Path:      "input.txt",
			WantPart1: "817",
			WantPart2: "925",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
