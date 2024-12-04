package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "21",
			WantPart2: "8",
		},
		{
			Path:      "input.txt",
			WantPart1: "1533",
			WantPart2: "345744",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
