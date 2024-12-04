package main

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	tests := aoc.Tests{Inputs: []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "13",
			WantPart2: "30",
		},
		{
			Path:      "input.txt",
			WantPart1: "23941",
			WantPart2: "5571760",
		},
	}}
	aoc.RunTests(t, tests, solvers1, solvers2)
}

func TestMain(m *testing.M) {
	log.SetFlags(0)
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}
