package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

const (
	borderCh   = ' '
	guardCh    = '^'
	obstacleCh = '#'
	freeCh     = '.'
)

var dirs = []fld.Pos{fld.Up, fld.Right, fld.Down, fld.Left}

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	field = field.AddBorder(borderCh)
	guardPos := field.FindFirst(guardCh)
	return len(visitedPositionsUntilGone(field, guardPos))
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	field = field.AddBorder(borderCh)
	guardPos := field.FindFirst(guardCh)

	ans := 0
	// Iterate over visited in part1 positions only: speeds up ~4 times.
	for curPos := range visitedPositionsUntilGone(field, guardPos) {
		if field.Get(curPos) != freeCh {
			continue
		}
		field.Set(curPos, obstacleCh)
		ans += aoc.BoolToInt(isLooped(field, guardPos))
		field.Set(curPos, freeCh)
	}

	return ans
}

func visitedPositionsUntilGone(field fld.ByteField, guardPos fld.Pos) containers.Set[fld.Pos] {
	dirIdx := 0
	visited := containers.NewSet[fld.Pos]()
	for field.Get(guardPos) != borderCh {
		npos := guardPos.Add(dirs[dirIdx])
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
			continue
		}
		visited.Add(guardPos)
		guardPos = npos
	}
	return visited
}

func isLooped(field fld.ByteField, guardPos fld.Pos) bool {
	dirIdx := 0
	type state struct {
		pos    fld.Pos
		dirIdx int
	}
	seen := containers.NewSet[state]()
	for field.Get(guardPos) != borderCh {
		curState := state{pos: guardPos, dirIdx: dirIdx}
		if seen.Has(curState) {
			return true
		}
		npos := guardPos.Add(dirs[dirIdx])
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
			continue
		}
		seen.Add(curState)
		guardPos = npos
	}
	return false
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
