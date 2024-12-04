package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{Path: "input_ex1.txt", WantPart1: "114", WantPart2: "2"},
		{Path: "input.txt", WantPart1: "1647269739", WantPart2: "864"},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}
