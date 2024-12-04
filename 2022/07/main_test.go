package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "95437",
			WantPart2: "24933642",
		},
		{
			Path:      "input.txt",
			WantPart1: "1454188",
			WantPart2: "4183246",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
