package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "405",
			WantPart2: "400",
		},
		{
			Path:      "input.txt",
			WantPart1: "37025",
			WantPart2: "32854",
		},
	}}
	aoc.RunTests(t, tests, solversPart1, solversPart2)
}
