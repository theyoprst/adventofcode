package main

import (
	"iter"
	"strings"

	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

func SolvePart1(lines []string) any {
	blocks := aoc.Blocks(lines)
	field := fld.NewByteField(blocks[0])
	robotPos := field.FindFirst('@')
	for dir := range parseCommands(blocks[1]) {
		// find first no box space in the direction dir:
		noBoxPos := robotPos.Add(dir)
		for field.Get(noBoxPos) == 'O' {
			noBoxPos = noBoxPos.Add(dir)
		}
		if field.Get(noBoxPos) == '.' {
			field.Set(noBoxPos, 'O')
			field.Set(robotPos, '.')
			field.Set(robotPos.Add(dir), '@')
			robotPos = robotPos.Add(dir)
		}
	}
	return score(field, 'O')
}

func SolvePart2(lines []string) any {
	blocks := aoc.Blocks(lines)
	field := fld.NewByteField(enlarge(blocks[0]))
	robotPos := field.FindFirst('@')

	// toPush returns list of items to push is topological reverse sorted order (the further item is the first in the list).
	// toPush returns empty list if there is a wall in the way.
	toPush := func(dir fld.Pos) []fld.Pos {
		visited := containers.NewSet[fld.Pos]()
		var topRevSorted []fld.Pos
		canPush := true

		var dfs func(from fld.Pos)
		dfs = func(from fld.Pos) {
			if visited.Has(from) {
				return
			}
			visited.Add(from)
			next := from.Add(dir)
			switch field.Get(next) {
			case '[':
				dfs(next)
				dfs(next.Add(fld.Right))
			case ']':
				dfs(next)
				dfs(next.Add(fld.Left))
			case '#':
				canPush = false
			}
			topRevSorted = append(topRevSorted, from)
		}

		dfs(robotPos)
		if !canPush {
			return nil
		}
		return topRevSorted
	}

	for dir := range parseCommands(blocks[1]) {
		for _, src := range toPush(dir) {
			// Because items to push are sorted (the further item is the first in the list),
			// there is no overlapping during pushing.
			field.Swap(src, src.Add(dir))
			if src == robotPos {
				robotPos = robotPos.Add(dir)
			}
		}
	}

	return score(field, '[')
}

func parseCommands(lines []string) iter.Seq[fld.Pos] {
	return func(yield func(fld.Pos) bool) {
		for _, command := range strings.Join(lines, "") {
			var dir fld.Pos
			switch command {
			case '^':
				dir = fld.Up
			case 'v':
				dir = fld.Down
			case '<':
				dir = fld.Left
			case '>':
				dir = fld.Right
			default:
				panic("invalid command")
			}
			if !yield(dir) {
				break
			}
		}
	}
}

func enlarge(lines []string) []string {
	var newLines []string
	for i := range lines {
		var sb strings.Builder
		for j := range len(lines[i]) {
			switch lines[i][j] {
			case '#':
				sb.WriteString("##")
			case 'O':
				sb.WriteString("[]")
			case '.':
				sb.WriteString("..")
			case '@':
				sb.WriteString("@.")
			}
		}
		newLines = append(newLines, sb.String())
	}
	return newLines
}

func score(field fld.ByteField, boxCh byte) int {
	sum := 0
	for pos := range field.IterPositions() {
		if field.Get(pos) == boxCh {
			sum += 100*pos.Row + pos.Col
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
