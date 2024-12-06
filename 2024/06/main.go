package main

import (
	"github.com/theyoprst/adventofcode/aoc"
	"github.com/theyoprst/adventofcode/aoc/containers"
	"github.com/theyoprst/adventofcode/aoc/fld"
)

const (
	guardCh    = '^'
	obstacleCh = '#'
	freeCh     = '.'
)

var dirs = []fld.Pos{fld.Up, fld.Right, fld.Down, fld.Left}

func SolvePart1(lines []string) any {
	field := fld.NewByteField(lines)
	guardPos := field.FindFirst(guardCh)
	return len(visitedPositionsUntilGone(field, guardPos))
}

func SolvePart2(lines []string) any {
	field := fld.NewByteField(lines)
	guardPos := field.FindFirst(guardCh)

	ans := 0
	// Iterate over visited in part1 positions only: speeds up ~4 times.
	for curPos, dirIdx := range visitedPositionsAndDirsUntilGone(field, guardPos) {
		if field.Get(curPos) != freeCh {
			continue
		}
		field.Set(curPos, obstacleCh)

		prevPos := curPos.Add(dirs[dirIdx].Reverse())
		ans += aoc.BoolToInt(isLooped(field, prevPos, dirIdx))

		field.Set(curPos, freeCh)
	}

	return ans
}

func visitedPositionsUntilGone(field fld.ByteField, guardPos fld.Pos) containers.Set[fld.Pos] {
	dirIdx := 0
	visited := containers.NewSet(guardPos)
	for {
		npos := guardPos.Add(dirs[dirIdx])
		if !field.Inside(npos) {
			break
		}
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
		} else {
			guardPos = npos
			visited.Add(guardPos)
		}
	}
	return visited
}

func visitedPositionsAndDirsUntilGone(field fld.ByteField, guardPos fld.Pos) map[fld.Pos]int {
	dirIdx := 0
	visited := make(map[fld.Pos]int)
	visited[guardPos] = dirIdx
	for {
		npos := guardPos.Add(dirs[dirIdx])
		if !field.Inside(npos) {
			break
		}
		if field.Get(npos) == obstacleCh {
			dirIdx = (dirIdx + 1) % len(dirs) // Turn right.
		} else {
			if _, ok := visited[npos]; !ok {
				visited[npos] = dirIdx
			}
			guardPos = npos
		}
	}
	return visited
}

func isLooped(field fld.ByteField, guardPos fld.Pos, dirIdx int) bool {
	type State struct {
		pos    fld.Pos
		dirIdx int
	}
	state := State{pos: guardPos, dirIdx: dirIdx}
	seen := containers.NewSet[State]()
	for !seen.Has(state) {
		npos := state.pos.Add(dirs[state.dirIdx])
		if !field.Inside(npos) {
			return false
		}
		if field.Get(npos) == obstacleCh {
			seen.Add(state)                               // Remember the state only on turn (speeds up 2.5-3 times).
			state.dirIdx = (state.dirIdx + 1) % len(dirs) // Turn right.
		} else {
			state.pos = npos
		}
	}
	return true
}

var (
	solvers1 = []aoc.Solver{SolvePart1}
	solvers2 = []aoc.Solver{SolvePart2}
)

func main() {
	aoc.Main(solvers1, solvers2)
}
