package main

import (
	"context"
	"slices"

	"github.com/theyoprst/adventofcode/aoc"
)

func SolvePart1(_ context.Context, lines []string) any {
	var locks, keys [][]int
	for _, block := range aoc.Blocks(lines) {
		hist := slices.Repeat([]int{-1}, len(block[0]))
		for _, row := range block {
			for i, c := range row {
				if c == '#' {
					hist[i]++
				}
			}
		}
		if block[0][0] == '#' {
			locks = append(locks, hist)
		} else {
			keys = append(keys, hist)
		}
	}

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if match(lock, key) {
				count++
			}
		}
	}
	return count
}

func match(lock, key []int) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
