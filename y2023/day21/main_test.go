package main

import (
	"fmt"
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "42",
		},
		{
			Path:      "input.txt",
			WantPart1: "3632",
			WantPart2: "600336060511101",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}

func TestInf3x3(t *testing.T) {
	lines := []string{
		"...",
		".S.",
		"...",
	}
	size := len(lines)
	for tiles := 2; tiles <= 10; tiles += 2 {
		steps := size/2 + tiles*size
		t.Run(fmt.Sprintf("steps=%d", steps), func(t *testing.T) {
			naiveAns := CountReachableInfiniteNaive(lines, steps)
			smartAns := CountReachableInfiniteSmart(lines, steps)
			if naiveAns != smartAns {
				t.Errorf("Answers mismatch: naive: %d, smart: %d", naiveAns, smartAns)
			}
		})
	}
}
