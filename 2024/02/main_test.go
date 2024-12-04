package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "2",
			WantPart2: "4",
		},
		{
			Path:      "input.txt",
			WantPart1: "224",
			WantPart2: "293",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
