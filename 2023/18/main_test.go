package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "62",
			WantPart2: "952408144115",
		},
		{
			Path:      "input.txt",
			WantPart1: "42317",
			WantPart2: "83605563360288",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
