package main

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

func Test(t *testing.T) {
	inputs := []aoc.Input{
		{
			Path:      "input_ex1.txt",
			WantPart1: "35",
			WantPart2: "46",
		},
		{
			Path:      "input.txt",
			WantPart1: "175622908",
			WantPart2: "5200543",
		},
	}
	aoc.RunTests(t, inputs, solvers1, solvers2)
}

func TestMain(m *testing.M) {
	log.SetFlags(0)
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}
