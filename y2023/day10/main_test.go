package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/theyoprst/adventofcode/aoc"
)

type WantFile struct {
	Part1 string `yaml:"part1"`
	Part2 string `yaml:"part2"`
}

func Test(t *testing.T) {
	dir, err := os.Open(".")
	if err != nil {
		t.Fatal(err)
	}
	files, err := dir.ReadDir(-1)
	if err != nil {
		t.Fatal(err)
	}
	ff := map[string]struct{}{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		ff[f.Name()] = struct{}{}
	}
	const wantSuf = ".want.yaml"
	const testSuf = ".txt"
	type testCase struct {
		path string
		want string
	}
	var testCasesPart1 []testCase
	var testCasesPart2 []testCase
	for _, f := range aoc.MapSortedKeys(ff) {
		if !strings.HasSuffix(f, wantSuf) {
			continue
		}
		testF := f[:len(f)-len(wantSuf)] + testSuf
		if _, exist := ff[testF]; !exist {
			t.Errorf("No file %q found", testF)
		}
		wantData, err := os.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		var want WantFile
		if err := yaml.Unmarshal(wantData, &want); err != nil {
			t.Error(err)
		}
		if want.Part1 != "" {
			testCasesPart1 = append(testCasesPart1, testCase{
				path: testF,
				want: want.Part1,
			})
		}
		if want.Part2 != "" {
			testCasesPart2 = append(testCasesPart2, testCase{
				path: testF,
				want: want.Part2,
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
