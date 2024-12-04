package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{Path: "input_ex1.txt", WantPart1: "21", WantPart2: "525152"},
		{Path: "input.txt", WantPart1: "7599", WantPart2: "15454556629917"},
	}}
	aoc.RunTests(t, tests, solversPart1, solversPart2)
}
