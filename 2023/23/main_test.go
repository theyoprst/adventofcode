package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "94",
			WantPart2: "154",
		},
		{
			Path:      "input.txt",
			WantPart1: "2018",
			WantPart2: "6406",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
