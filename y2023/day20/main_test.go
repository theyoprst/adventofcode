package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "32000000",
		},
		{
			Path:      "input_ex2.txt",
			WantPart1: "11687500",
		},
		{
			Path:      "input.txt",
			WantPart1: "836127690",
			WantPart2: "240914003753369",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
