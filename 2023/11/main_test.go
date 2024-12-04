package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{Path: "input_ex1.txt", WantPart1: "374", WantPart2: "82000210"},
		{Path: "input.txt", WantPart1: "9522407", WantPart2: "544723432977"},
	}}
	aoc.RunTests(t, tests, solversPart1, solversPart2)
}
