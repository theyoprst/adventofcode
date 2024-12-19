package main

import (
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(lines []string) any {
	patterns := strings.Split(lines[0], ", ")
	designs := lines[2:]
	totalCount := 0
	for _, design := range designs {
		prefixIsDecomposable := make([]bool, len(design)+1)
		prefixIsDecomposable[0] = true
		for checkLen := 1; checkLen <= len(design); checkLen++ {
			for _, pattern := range patterns {
				if checkLen < len(pattern) {
					continue
				}
				if !prefixIsDecomposable[checkLen-len(pattern)] {
					continue
				}
				if design[checkLen-len(pattern):checkLen] == pattern {
					prefixIsDecomposable[checkLen] = true
					break
				}
			}
		}
		if prefixIsDecomposable[len(design)] {
			totalCount++
		}
	}
	return totalCount
}

func SolvePart2(lines []string) any {
	patterns := strings.Split(lines[0], ", ")
	designs := lines[2:]
	totalCount := 0
	for _, design := range designs {
		decomposePrefixCount := make([]int, len(design)+1)
		decomposePrefixCount[0] = 1
		for checkLen := 1; checkLen <= len(design); checkLen++ {
			for _, pattern := range patterns {
				if checkLen < len(pattern) {
					continue
				}
				if design[checkLen-len(pattern):checkLen] == pattern {
					decomposePrefixCount[checkLen] += decomposePrefixCount[checkLen-len(pattern)]
				}
			}
		}
		totalCount += decomposePrefixCount[len(design)]
	}
	return totalCount
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
