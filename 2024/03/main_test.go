package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "161",
			WantPart2: "161",
		},
		{
			Path:      "input_ex2.txt",
			WantPart1: "161",
			WantPart2: "48",
		},
		{
			Path:      "input.txt",
			WantPart1: "178538786",
			WantPart2: "102467299",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
