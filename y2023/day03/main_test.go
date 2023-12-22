package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{Path: "input_t1.txt", WantPart1: "4361", WantPart2: "467835"},
		{Path: "input.txt", WantPart1: "544664", WantPart2: "84495585"},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
