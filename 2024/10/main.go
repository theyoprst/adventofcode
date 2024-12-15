package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	score := func(pos fld.Pos) int {
		// BFS over the field starting with pos.
		// Returns number of '9's found.
		visited := make(map[fld.Pos]bool)
		visited[pos] = true
		queue := []fld.Pos{pos}
		ninesFound := 0
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			if field.Get(pos) == '9' {
				ninesFound++
			}
			for _, dir := range fld.DirsSimple {
				next := pos.Add(dir)
				if !field.Inside(next) || visited[next] || field.Get(next) != field.Get(pos)+1 {
					continue
				}
				visited[next] = true
				queue = append(queue, next)
			}
		}
		return ninesFound
	}

	sum := 0
	for pos := range field.IterPositions() {
		if field.Get(pos) == '0' {
			sum += score(pos)
		}
	}
	return sum
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	rating := func(pos fld.Pos) int {
		// Same BFS as in part 1, but with counting all unique paths.
		uniquePaths := make(map[fld.Pos]int)
		uniquePaths[pos] = 1
		queue := []fld.Pos{pos}
		totalPaths := 0
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			if field.Get(pos) == '9' {
				totalPaths += uniquePaths[pos]
			}
			for _, dir := range fld.DirsSimple {
				next := pos.Add(dir)
				if !field.Inside(next) || field.Get(next) != field.Get(pos)+1 {
					continue
				}
				if uniquePaths[next] == 0 {
					queue = append(queue, next)
				}
				uniquePaths[next] += uniquePaths[pos]
			}
		}
		return totalPaths
	}

	sum := 0
	for pos := range field.IterPositions() {
		if field.Get(pos) == '0' {
			sum += rating(pos)
		}
	}
	return sum
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
