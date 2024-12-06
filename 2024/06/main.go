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
	return len(visited)
}

func SolvePart2(lines []string) any {
	// Brute force. Put obstacles on the path and check if the guard will loop.
	// Optimizations:
	// 1. Put obstacles only on the path from part1, not on the whole field (speedup x4)
	// 2. Remember visited states only on turns, not on every step (speedup x3)
	// 3. Start loop detection from the new obstacle, not from the initial guard position (speedup x3)
	// Total: 0.13s for the input (from 5s without optimizations).
	// TODO: implement jumps to the obstacle.
	field := fld.NewByteField(lines)
	guardPos := field.FindFirst(guardCh)

	ans := 0
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
			if !visited.Has(npos) {
				field.Set(npos, obstacleCh)
				if isLooped(field, guardPos, dirIdx) {
					ans++
				}
				field.Set(npos, freeCh)
				visited.Add(npos)
			}
			guardPos = npos
		}
	}
	return ans
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
			seen.Add(state) // Remember the state only on turn (speeds up 3 times).
			state.dirIdx = (state.dirIdx + 1) % len(dirs)
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
