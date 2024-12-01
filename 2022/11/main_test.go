package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "10605",
			WantPart2: "2713310158",
		},
		{
			Path:      "input.txt",
			WantPart1: "56350",
			WantPart2: "13954061248",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
