package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "288",
			WantPart2: "71503",
		},
		{
			Path:      "input.txt",
			WantPart1: "1624896",
			WantPart2: "32583852",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
