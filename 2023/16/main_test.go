package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "46",
			WantPart2: "51",
		},
		{
			Path:      "input.txt",
			WantPart1: "6795",
			WantPart2: "7154",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
