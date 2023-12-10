package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/theyoprst/adventofcode/aoc"
)

type Input struct {
	Path      string
	WantPart1 string
	WantPart2 string
}

var inputs []Input = []Input{
	{Path: "input_ex0.txt", WantPart1: "4", WantPart2: "1"},
	{Path: "input_ex1.txt", WantPart1: "4", WantPart2: "1"},
	{Path: "input_ex2.txt", WantPart1: "8", WantPart2: "1"},
	{Path: "input_ex3.txt", WantPart2: "4"},
	{Path: "input_ex4.txt", WantPart2: "8"},
	{Path: "input.txt", WantPart1: "6897", WantPart2: "367"},
}

func Test(t *testing.T) {
	type testCase struct {
		path string
		want string
	}
	var testCasesPart1 []testCase
	var testCasesPart2 []testCase
	for _, input := range inputs {
		if input.WantPart1 != "" {
			testCasesPart1 = append(testCasesPart1, testCase{
				path: input.Path,
				want: input.WantPart1,
			})
		}
		if input.WantPart2 != "" {
			testCasesPart2 = append(testCasesPart2, testCase{
				path: input.Path,
				want: input.WantPart2,
			})
		}
	}

	type testSuite struct {
		name      string
		solver    func([]string) any
		testCases []testCase
	}

	testSuites := []testSuite{
		{
			name:      "Part1",
			solver:    SolvePart1,
			testCases: testCasesPart1,
		},
		{
			name:      "Part2",
			solver:    SolvePart2,
			testCases: testCasesPart2,
		},
	}

	for _, suite := range testSuites {
		t.Run(suite.name, func(t *testing.T) {
			for _, testCase := range suite.testCases {
				t.Run(testCase.path, func(t *testing.T) {
					f, err := os.Open(testCase.path)
					if err != nil {
						t.Fatal(err)
					}
					defer func() {
						_ = f.Close()
					}()
					ans := suite.solver(aoc.ReadLines(f))
					ansStr := fmt.Sprint(ans)
					if ansStr != testCase.want {
						t.Errorf("Got answer %q, want %q", ansStr, testCase.want)
					}
				})
			}
		})
	}
}
