package main

import (
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		// {
		// 	Path:      "input_ex1.txt",
		// 	WantPart1: "0",
		// 	WantPart2: "0",
		// },
		// {
		// 	Path:      "input.txt",
		// 	WantPart1: "0",
		// 	WantPart2: "0",
		// },
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}
