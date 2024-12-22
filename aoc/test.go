package aoc

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v3"
)

type Input struct {
	Path      string `yaml:"path"`
	WantPart1 string `yaml:"wantPart1"`
	WantPart2 string `yaml:"wantPart2"`
	Params    Params `yaml:"params"`
}

type Tests struct {
	Inputs []Input `yaml:"inputs"`
}

func RunTests(t *testing.T, solversPart1 []Solver, solversPart2 []Solver) {
	t.Helper()

	data, err := os.ReadFile("tests.yaml")
	if err != nil {
		t.Fatal(err)
	}

	var tests Tests
	if err := yaml.Unmarshal(data, &tests); err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		path   string
		want   string
		params Params
	}
	var testCasesPart1 []testCase
	var testCasesPart2 []testCase
	for _, input := range tests.Inputs {
		if input.WantPart1 != "" {
			testCasesPart1 = append(testCasesPart1, testCase{
				path:   input.Path,
				want:   input.WantPart1,
				params: input.Params,
			})
		}
		if input.WantPart2 != "" {
			testCasesPart2 = append(testCasesPart2, testCase{
				path:   input.Path,
				want:   input.WantPart2,
				params: input.Params,
			})
		}
	}

	type testSuite struct {
		name      string
		solver    func(context.Context, []string) any
		testCases []testCase
	}

	testSuites := make([]testSuite, 0, len(solversPart1)+len(solversPart1))
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
					ctx := context.Background()
					if testCase.params != nil {
						ctx = contextWithParams(ctx, testCase.params)
					}
					ans := suite.solver(ctx, ReadLines(f))
					ansStr := fmt.Sprint(ans)
					if diff := cmp.Diff(ansStr, testCase.want); diff != "" {
						t.Errorf("Unexpected answer, diff (-got +want):\n%s", diff)
					}
				})
			}
		})
	}
}
