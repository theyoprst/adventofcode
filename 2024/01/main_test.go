package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "11",
			WantPart2: "31",
		},
		{
			Path:      "input.txt",
			WantPart1: "2904518",
			WantPart2: "18650129",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
