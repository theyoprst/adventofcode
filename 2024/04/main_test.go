package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "18",
			WantPart2: "9",
		},
		{
			Path:      "input.txt",
			WantPart1: "2718",
			WantPart2: "2046",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
