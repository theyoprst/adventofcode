package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	aoc.RunTests(t, solvers1, solvers2)
}

func TestCmpWithNaive(t *testing.T) {
	testCases := []struct {
		path     string
		maxTiles int
	}{
		{path: "input_ex2.txt", maxTiles: 10},
		{path: "input.txt", maxTiles: 2},
	}
	for _, test := range testCases {
		t.Run(test.path, func(t *testing.T) {
			f, err := os.Open(test.path)
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				_ = f.Close()
			}()
			lines := aoc.ReadLines(f)
			size := len(lines)
			for tiles := 1; tiles <= test.maxTiles; tiles++ {
				steps := size/2 + tiles*size
				t.Run(fmt.Sprintf("steps=%d", steps), func(t *testing.T) {
					naiveAns := CountReachableInfiniteNaive(lines, steps)
					smartAns := CountReachableInfiniteSmart(lines, steps)
					quadraticAns := CountReachableInfiniteQuadratic(lines, steps)
					if naiveAns != smartAns {
						t.Errorf("Answers mismatch: naive: %d, smart: %d", naiveAns, smartAns)
					}
					if naiveAns != quadraticAns {
						t.Errorf("Answers mismatch: naive: %d, quadratic: %d", naiveAns, quadraticAns)
					}
				})
			}
		})
	}
}
