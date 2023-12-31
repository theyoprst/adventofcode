package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "19114",
			WantPart2: "167409079868000",
		},
		{
			Path:      "input.txt",
			WantPart1: "373302",
			WantPart2: "130262715574114",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
