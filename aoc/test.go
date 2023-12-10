package aoc

import (
	"fmt"
	"os"
	"testing"
)

type Input struct {
	Path      string
	WantPart1 string
	WantPart2 string
}

func RunTests(t *testing.T, inputs []Input, solversPart1 []Solver, solversPart2 []Solver) {
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

	var testSuites []testSuite
	for _, solvePart1 := range solversPart1 {
		testSuites = append(testSuites, testSuite{
			name:      getFunctionName(solvePart1),
			solver:    solvePart1,
			testCases: testCasesPart1,
		})
	}
	for _, solvePart2 := range solversPart2 {
		testSuites = append(testSuites, testSuite{
			name:      getFunctionName(solvePart2),
			solver:    solvePart2,
			testCases: testCasesPart2,
		})
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
					ans := suite.solver(ReadLines(f))
					ansStr := fmt.Sprint(ans)
					if ansStr != testCase.want {
						t.Errorf("Got answer %q, want %q", ansStr, testCase.want)
					}
				})
			}
		})
	}
}
