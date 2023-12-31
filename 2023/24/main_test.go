package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "0",
			WantPart2: "47",
		},
		{
			Path:      "input.txt",
			WantPart1: "14799",
			WantPart2: "1007148211789625",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
