package main

import (
	"context"
	"fmt"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
	"github.com/theyoprst/adventofcode/must"
)

func SolvePart1(_ context.Context, lines []string) any {
	return solve(lines, 2)
}

func SolvePart2(_ context.Context, lines []string) any {
	return solve(lines, 20)
}

func solve(lines []string, cheatSize int) int {
	field := fld.NewByteField(lines)
	dist := calculateDistances(field)
	profits := make(map[int]int)
	for pos := range field.IterPositionsWithPadding(1) {
		if field.Get(pos) == '#' {
			continue
		}
		for dRow := -cheatSize; dRow <= cheatSize; dRow++ {
			for dCol := -cheatSize; dCol <= cheatSize; dCol++ {
				jumpSize := aoc.Abs(dRow) + aoc.Abs(dCol)
				if jumpSize > cheatSize {
					continue
				}
				next := fld.Pos{Row: pos.Row + dRow, Col: pos.Col + dCol}
				if !field.Inside(next) {
					continue
				}
				if field.Get(next) == '#' {
					continue
				}
				profit := dist[next] - dist[pos] - jumpSize
				if profit > 0 {
					profits[profit]++
				}
			}
		}
	}

	const threshold = 100
	totalProfits := 0
	for profit, count := range profits {
		if profit >= threshold {
			totalProfits += count
		}
	}

	// Debug for part1
	// fmt.Println(profits)
	return totalProfits
}

func calculateDistances(field fld.ByteField) map[fld.Pos]int {
	start := field.FindFirst('S')
	finish := field.FindFirst('E')
	dist := make(map[fld.Pos]int)
	dist[start] = 0
	for pos := start; pos != finish; {
		var next fld.Pos
		for _, dir := range fld.DirsSimple {
			cand := pos.Add(dir)
			if _, candVisited := dist[cand]; !candVisited && field.Get(cand) != '#' {
				if next != fld.Zero {
					must.NoError(fmt.Errorf("several paths from %v", pos))
				}
				next = cand
			}
		}
		dist[next] = dist[pos] + 1
		pos = next
	}
	return dist
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
